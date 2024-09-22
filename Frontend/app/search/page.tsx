"use client";

import { ChangeEvent, useState } from "react";
import { NFT, Table_Basic, TagType } from "../types";
import { useRouter, useSearchParams } from "next/navigation";
import { AiTag } from "../components/CustomTag";
import Pagination from "../components/Pagination";
import Loading from "../components/Loading";
import { useAsyncEffect, useDebounceEffect } from "ahooks";
import Link from "next/link";

const types: TagType[] = [
  { name: "最新", id: 0 },
  { name: "最热", id: 1 },
  { name: "畅销", id: 2 },
];
const PageSize = 20;
export default function Page() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [tableData, setTableData] = useState<NFT[]>([]);
  const [currentType, setCurrentType] = useState<TagType>(types[0]);
  const [currentPage, setCurrentPage] = useState<number>(1);
  const [total, setTotal] = useState<number>(0);
  const [loading, setLoading] = useState(false);
  const [selectedSubject, setSelectSubject] = useState<number[]>([]);
  const [selectedTopic, setSelectTopic] = useState<number[]>([]);

  const [currentGrade, setCurrentGrade] = useState<number | null>();
  const [gradeList, setGradeList] = useState<TagType[]>([]);

  const [subjectList, setSubjectList] = useState<TagType[]>([]);

  const [topicList, setTopicList] = useState<TagType[]>([]);

  const onSwitchType = (item: TagType) => {
    setCurrentType(item);
  };

  const getSubjectList = async () => {
    if (currentGrade) {
      const response = await (
        await fetch(`/api/subject?id=${currentGrade}`)
      ).json();

      setSelectSubject([]);
      setSelectTopic([]);
      if (response.count > 0) {
        setSubjectList(response.data);
      }
    }
  };
  const getTopicList = async () => {
    if (currentGrade) {
      const response = await (
        await fetch(
          `/api/topic?gradeId=${currentGrade}&subjectId=${selectedSubject.join(",")}`
        )
      ).json();
      setSelectTopic([]);
      if (response.count > 0) {
        setTopicList(response.data);
      }
    }
  };

  useAsyncEffect(async () => {
    const response = await (await fetch("/api/grade")).json();
    if (response.count > 0) {
      setGradeList(response.data);
      setCurrentGrade(response.data[0].id);
    }
  }, []);

  useDebounceEffect(() => {
    getSubjectList();
  }, [currentGrade]);

  useDebounceEffect(() => {
    getTopicList();
  }, [currentGrade, selectedSubject]);

  const onChangeGrade = (grade: TagType) => {
    setCurrentGrade(grade.id);
  };

  useDebounceEffect(
    () => {
      //
    },
    [currentGrade, selectedSubject],
    {
      wait: 2000,
    }
  );
  const onSearch = async () => {
    const response = await fetch("/api/search", {
      method: "POST",
      body: JSON.stringify({
        keyword: searchParams.get("key"),
        gradeIds: [currentGrade],
        subjectIds: selectedSubject,
        topicIds: selectedTopic,
        page: 1,
        pagesize: PageSize,
      }),
    });
    const responseJson = await response.json();
    if (!!responseJson.count) {
      setTableData(responseJson.data);
      setTotal(responseJson.count);
    } else {
      setTableData([]);
      setTotal(0);
    }
  };

  const onSubjectChange = (checked: boolean, id: number) => {
    const selected = [...selectedSubject];
    if (checked) selected.push(id);
    else {
      selected.splice(selected.indexOf(id), 1);
    }
    setCurrentPage(1);
    setSelectSubject(selected);
  };
  const onTopicChange = (checked: boolean, id: number) => {
    const selected = [...selectedTopic];
    if (checked) selected.push(id);
    else {
      selected.splice(selected.indexOf(id), 1);
    }
    setSelectTopic(selected);
  };

  useDebounceEffect(
    () => {
      onSearch();
    },
    [currentPage, selectedSubject, selectedTopic],
    {
      wait: 2000,
    }
  );
  return (
    <div>
      <Loading show={loading} />
      <div className="flex items-start gap-4">
        {/* left start */}
        <div className="w-[332px]">
          <div className="w-full min-h-80 bg-white rounded-sm shadow-sm mb-4">
            <div className="py-4 text-center bg-gray-50">学科分类</div>
            <div className="p-4">
              <div>
                {gradeList.map((item) => (
                  <span
                    key={item.id}
                    className={`${currentGrade === item.id ? "text-white bg-primary" : "hover:bg-gray-100"} inline-block py-1 px-4 rounded-sm text-sm font-light text-[#666] cursor-pointer`}
                    onClick={() => onChangeGrade(item)}
                  >
                    {item.name}
                  </span>
                ))}
              </div>
              <div className="mt-4">
                {subjectList.map((item) => (
                  <div className="my-3 flex items-center">
                    <input
                      type="checkbox"
                      onChange={(e: ChangeEvent<HTMLInputElement>) =>
                        onSubjectChange(e.target.checked, item.id)
                      }
                      className="form-checkbox mr-1 rounded-sm w-4 h-4 border-[#979797] focus:ring-offset-0 focus:ring-0"
                    />
                    <label className="leading-none text-[#666]">
                      {item.name}
                    </label>
                  </div>
                ))}
              </div>
            </div>
          </div>
          <div className="min-h-80 bg-white pb-14 rounded-sm shadow-sm">
            <div className="py-4 text-center bg-gray-50">主题分类</div>
            <div className="px-4 pt-2 pb-4">
              {topicList.map((item) => (
                <div className="my-3 flex items-center">
                  <input
                    type="checkbox"
                    onChange={(e: ChangeEvent<HTMLInputElement>) =>
                      onTopicChange(e.target.checked, item.id)
                    }
                    className="form-checkbox mr-1 rounded-sm w-4 h-4 border-[#979797] focus:ring-offset-0 focus:ring-0"
                  />
                  <label className="leading-none text-[#666]">
                    {item.name}
                  </label>
                </div>
              ))}
            </div>
          </div>
        </div>
        {/* left end */}
        {/* right start */}
        <div className="flex-1 min-h-[41rem] relative bg-white rounded-sm shadow-sm">
          <div className="py-3 px-4 h-14 flex justify-between items-center bg-gray-50">
            <p className="font-light text-xs text-[#666]">
              当前总有<span className="text-primary text-base">{total}</span>
              个作品
            </p>
            <ul className="flex w-36 justify-around bg-gray-100 rounded-md *:text-sm *:hover:cursor-pointer">
              {/* {types.map((item) => (
                <li
                  className={`w-full py-1 rounded-md text-center ${currentType.id === item.id ? " bg-blue-200" : " bg-gray-100"}`}
                  onClick={() => onSwitchType(item)}
                >
                  {item.name}
                </li>
              ))} */}
            </ul>
          </div>
          <div className="p-4">
            <table className="w-full">
              <thead>
                <tr className="*:px-2 *:py-3 *:text-left *:text-xs *:font-normal *:text-primary-font_9 *:uppercase">
                  <th className="w-12">序号</th>
                  <th className="min-w-36">创意名称</th>
                  <th className="w-24">作者</th>
                  <th className="w-24">发布日期</th>
                  <th className="w-24">售价</th>
                  <th className="w-24">热度</th>
                  <th className="w-24">销量</th>
                  {/* <th className="w-24">被引用次数</th> */}
                </tr>
              </thead>
              <tbody className="bg-white">
                {tableData.slice(0, 5).map((item, i) => (
                  <tr className="*:p-2 *:whitespace-nowrap overflow-hidden cursor-pointer hover:bg-blue-50 rounded-md">
                    <td className="w-12">
                      {PageSize * (currentPage - 1) + i + 1}
                    </td>
                    <td className="min-w-36">
                      <Link href={`/nft/${item.ID}`}>{item.Title}</Link>
                      {/* {item.isAi ? <AiTag /> : <></>} */}
                    </td>
                    <td className="w-24 text-blue-400">
                      {item.Creator.Username}
                    </td>
                    <td className="w-24">{item.CreatedAt}</td>
                    <td className="w-24">{item.Price}ETH</td>
                    <td className="w-24">{item.LikeCount}</td>
                    <td className="w-24">{item.TransactionCount}</td>
                    {/* <td className="w-24">{item.sellPrice}</td> */}
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
          <Pagination
            total={total}
            currentPage={currentPage}
            onPageChange={(page: number) => setCurrentPage(page)}
          />
        </div>
        {/* right end */}
      </div>
    </div>
  );
}
