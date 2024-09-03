"use client";

import { useState } from "react";
import { Table_Basic, TagType } from "../types";
import { useRouter } from "next/navigation";

const grades: TagType[] = [
  { name: "全部", key: -1 },
  { name: "小学", key: 0 },
  { name: "初中", key: 1 },
  { name: "高中", key: 2 },
  { name: "高校", key: 3 },
];

const subjects: TagType[] = [
  { name: "哲学", key: 0 },
  { name: "法学", key: 1 },
  { name: "经济学", key: 2 },
  { name: "文学", key: 3 },
  { name: "理学", key: 4 },
  { name: "教育学", key: 5 },
  { name: "历史学", key: 6 },
  { name: "工学", key: 7 },
  { name: "农学", key: 5 },
  { name: "医学", key: 6 },
  { name: "军事学", key: 7 },
];

const topics: TagType[] = [
  { name: "电子", key: 0 },
  { name: "航天", key: 1 },
  { name: "设计", key: 2 },
  { name: "机器学习", key: 3 },
  { name: "市场营销", key: 4 },
  { name: "古代文学", key: 5 },
  { name: "现代文学", key: 6 },
  { name: "中国近代史", key: 7 },
  { name: "美国近代史", key: 8 },
  { name: "植物学", key: 9 },
  { name: "军事学", key: 10 },
];

const tableData: Table_Basic[] = [
  {
    index: 1,
    name: "创意点子",
    publishDate: "2024-08-31",
    sellPrice: "1.725 ETH",
  },
  {
    index: 2,
    name: "创意点子",
    publishDate: "2024-08-31",
    sellPrice: "1.725 ETH",
  },
  {
    index: 3,
    name: "创意点子",
    publishDate: "2024-08-31",
    sellPrice: "1.725 ETH",
  },
  {
    index: 4,
    name: "创意点子",
    publishDate: "2024-08-31",
    sellPrice: "1.725 ETH",
  },
  {
    index: 5,
    name: "创意点子",
    publishDate: "2024-08-31",
    sellPrice: "1.725 ETH",
  },
  {
    index: 6,
    name: "创意点子",
    publishDate: "2024-08-31",
    sellPrice: "1.725 ETH",
  },
  {
    index: 7,
    name: "创意点子",
    publishDate: "2024-08-31",
    sellPrice: "1.725 ETH",
  },
];

const types: TagType[] = [
  { name: "最新", key: 0 },
  { name: "最热", key: 1 },
  { name: "畅销", key: 2 },
];

export default function Page() {
  const router = useRouter();
  const [currentGrade, setCurrentGrade] = useState(-1);
  const [currentType, setCurrentType] = useState<TagType>(types[0]);

  const onSwitchType = (item: TagType) => {
    setCurrentType(item);
  };

  const onChangeGrade = (grade: TagType) => {
    setCurrentGrade(grade.key);
  };
  return (
    <div className="flex items-start gap-4">
      {/* left start */}
      <div className="w-64">
        <div className="min-h-80 bg-white rounded-sm shadow-sm mb-4">
          <div className="py-3 text-center bg-gray-50">学科分类</div>
          <div className="p-4">
            <div>
              {grades.map((item) => (
                <span
                  key={item.key}
                  className={`${currentGrade === item.key ? "text-white bg-blue-500" : "hover:bg-gray-100"} inline-block py-0.5 px-2 mr-1 rounded-sm text-[0.7rem] font-light text-gray-500 cursor-pointer`}
                  onClick={() => onChangeGrade(item)}
                >
                  {item.name}
                </span>
              ))}
            </div>
            <div className="mt-4">
              {subjects.map((item) => (
                <div className="my-3 flex items-center">
                  <input
                    type="checkbox"
                    className="form-checkbox mr-1 rounded-sm w-[0.8rem] h-[0.8rem] border-gray-500 focus:ring-offset-0 focus:ring-0"
                  />
                  <label className="leading-none text-[0.rem] text-gray-500">
                    {item.name}
                  </label>
                </div>
              ))}
            </div>
          </div>
        </div>
        <div className="min-h-80 bg-white rounded-sm shadow-sm">
          <div className="py-3 text-center bg-gray-50">主题分类</div>
          <div className="px-4 pt-2 pb-4">
            {topics.map((item) => (
              <div className="my-3 flex items-center">
                <input
                  type="checkbox"
                  className="form-checkbox mr-1 rounded-sm w-[0.8rem] h-[0.8rem] border-gray-500 focus:ring-offset-0 focus:ring-0"
                />
                <label className="leading-none text-[0.rem] text-gray-500">
                  {item.name}
                </label>
              </div>
            ))}
          </div>
        </div>
      </div>
      {/* left end */}
      {/* right start */}
      <div className="flex-1 min-h-[41rem] bg-white rounded-sm shadow-sm">
        <div className="py-3 px-4 flex justify-between items-center bg-gray-50">
          <p className="font-light text-[0.7rem]">
            当前总有<span className="text-sm text-blue-400">12386</span>
            个作品
          </p>
          <ul className="flex w-36 justify-around bg-gray-100 rounded-md text-xs *:text-xs *:hover:cursor-pointer">
            {types.map((item) => (
              <li
                className={`w-full py-1 rounded-md text-center ${currentType.key === item.key ? " bg-blue-200" : " bg-gray-100"}`}
                onClick={() => onSwitchType(item)}
              >
                {item.name}
              </li>
            ))}
          </ul>
        </div>
        <div className="p-4">
          <table className="w-full">
            <thead>
              <tr className="*:px-2 *:py-3 *:text-left *:font-normal *:text-[0.7rem] *:text-gray-400 *:uppercase">
                <th className="w-10">序号</th>
                <th className="min-w-36">创意名称</th>
                <th className="w-24">作者</th>
                <th className="w-24">发布日期</th>
                <th className="w-24">售价</th>
                <th className="w-24">热度</th>
                <th className="w-24">销量</th>
                <th className="w-24">被引用次数</th>
              </tr>
            </thead>
            <tbody className="bg-white">
              {tableData.slice(0, 5).map((item) => (
                <tr
                  onClick={() => router.push(`/idea/${item.index}`)}
                  className="*:p-2 *:whitespace-nowrap *:text-xs overflow-hidden cursor-pointer hover:bg-blue-50 rounded-md"
                >
                  <td className="text-gray-400 w-10">{item.index}</td>
                  <td className="min-w-36">{item.name}</td>
                  <td className="w-24 text-blue-400">{item.name}</td>
                  <td className="w-24">{item.publishDate}</td>
                  <td className="w-24">{item.sellPrice}</td>
                  <td className="w-24">{item.name}</td>
                  <td className="w-24">{item.publishDate}</td>
                  <td className="w-24">{item.sellPrice}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
      {/* right end */}
    </div>
  );
}
