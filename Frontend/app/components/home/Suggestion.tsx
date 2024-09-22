"use client";

import { NFT, TagType } from "@/app/types";
import { useState } from "react";
import { useRouter } from "next/navigation";
import ChangeFollowTopic from "./ChangeFollowTopic";
// import Tab from "../Tab";
import { OrderTag } from "../CustomTag";
import { useAsyncEffect } from "ahooks";

const topic: TagType[] = [
  { name: "推荐", id: 0 },
  // { name: "电子信息", id: 1 },
  // { name: "航空航天", id: 2 },
  // { name: "人脸识别", id: 3 },
];

const types: TagType[] = [
  { name: "最新", id: 0 },
  { name: "最热", id: 1 },
  { name: "畅销", id: 2 },
];

export default function HomeSuggestion() {
  const router = useRouter();
  const [tableData, setTableData] = useState<NFT[]>([]);
  const [isAdd, setIsAdd] = useState<boolean>(false);
  const [topics, setTopics] = useState(topic);
  const [currentTopic, setCurrentTopic] = useState<TagType>(topic[0]);

  const onSwitchTopic = (item: TagType) => {
    setCurrentTopic(item);
  };
  const onChange = (item: TagType) => {
    //
  };

  const onUpdateFollowTopics = (newTopics: TagType[]) => {
    setTopics(newTopics.slice());
    setIsAdd(false);
  };

  useAsyncEffect(async () => {
    try {
      const response = await fetch("/api/search", {
        method: "POST",
        body: JSON.stringify({
          keyword: "",
          gradeIds: [],
          subjectIds: [],
          topicIds: [],
          page: 1,
          pagesize: 10,
        }),
      });
      const responseJson = await response.json();
      if (!!responseJson.count) {
        setTableData(responseJson.data);
      } else {
        setTableData([]);
      }
    } catch {
      //
    }
  }, []);

  return (
    <div className="w-full my-8 px-10 py-8 bg-white rounded border border-primary-border">
      {isAdd ? (
        <ChangeFollowTopic
          followTopics={topics}
          onUpdateFollowTopics={onUpdateFollowTopics}
        />
      ) : (
        <>
          {/* header start */}
          <div className="flex justify-between gap-2 items-center">
            <ul className="flex gap-2 *:w-20 *:text-center *:rounded-full *:px-2 *:py-1 *:text-[0.75rem] *:hover:cursor-pointer">
              {topics.map((item) => (
                <li
                  key={item.id}
                  className={`${currentTopic.id === item.id ? "text-white bg-primary" : "text-black bg-primary-light_bg/50"}`}
                  onClick={() => onSwitchTopic(item)}
                >
                  {item.name}
                </li>
              ))}
              {/* <li
                className="text-black bg-primary-light_bg/50"
                onClick={() => setIsAdd(true)}
              >
                + 新增主题
              </li> */}
            </ul>
            <div className="flex gap-2">
              {/* <Tab data={types} onChange={onChange} />
              <div className="bg-primary-light_bg/50 px-2 py-1 rounded-md text-[0.75rem] cursor-pointer">
                查看更多
              </div> */}
            </div>
          </div>
          {/* header end */}

          <div className="flex gap-6 my-4 min-h-0">
            <div className="w-1/2">
              <table className="w-full">
                <thead>
                  <tr className="*:px-2 *:py-3 *:text-left *:font-normal *:text-sm *:text-[#999] *:uppercase">
                    <th>序号</th>
                    <th>创意名称</th>
                    <th>发布日期</th>
                    <th>售价</th>
                  </tr>
                </thead>
                <tbody className="bg-white">
                  {tableData.slice(0, 5).map((item, i) => (
                    <tr
                      key={i}
                      onClick={() => router.push(`/nft/${item.ID}`)}
                      className="*:p-2 *:whitespace-nowrap *:text-[#333] overflow-hidden cursor-pointer hover:bg-blue-50 rounded-md"
                    >
                      <td>
                        <OrderTag order={i + 1} bg={i < 3} />
                      </td>
                      <td>{item.Title}</td>
                      <td>{item.CreatedAt}</td>
                      <td>{item.Price}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
            <div className="w-1/2">
              {tableData.length > 5 && (
                <table className="w-full">
                  <thead>
                    <tr className="*:px-2 *:py-3 *:text-left *:font-normal *:text-sm *:text-[#999] *:uppercase">
                      <th>序号</th>
                      <th>创意名称</th>
                      <th>发布日期</th>
                      <th>售价</th>
                    </tr>
                  </thead>
                  <tbody>
                    {tableData.slice(5).map((item, i) => (
                      <tr
                        key={i}
                        onClick={() => router.push(`/nft/${item.ID}`)}
                        className="*:p-2 *:whitespace-nowrap *:text-[#333] overflow-hidden cursor-pointer hover:bg-blue-50 rounded-md"
                      >
                        <td>
                          <OrderTag order={6 + i} bg={false} />
                        </td>
                        <td>{item.Title}</td>
                        <td>{item.CreatedAt}</td>
                        <td>{item.Price}</td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              )}
            </div>
          </div>
        </>
      )}
    </div>
  );
}
