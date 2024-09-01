"use client";
import { Table_Basic, TagType } from "@/app/types";
import { useState } from "react";

const types: TagType[] = [
  { name: "一天", key: 0 },
  { name: "一周", key: 1 },
  { name: "一月", key: 2 },
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

export default function HotSellerList() {
  const [currentType, setCurrentType] = useState<TagType>(types[0]);

  const onSwitchType = (item: TagType) => {
    setCurrentType(item);
  };
  return (
    <div className="w-full my-4 p-4 bg-white rounded shadow-sm">
      {/* header start */}
      <div className="flex justify-between gap-2 items-center">
        <div>热门榜单</div>
        <div className="flex gap-2">
          <ul className="flex w-36 justify-around bg-gray-100 rounded-md text-xs *:text-xs *:hover:cursor-pointer">
            {types.map((item) => (
              <li
                className={`w-full py-1 rounded-md text-center ${currentType.key === item.key ? " bg-gray-200" : " bg-gray-100"}`}
                onClick={() => onSwitchType(item)}
              >
                {item.name}
              </li>
            ))}
          </ul>
          <div className="bg-gray-100 px-2 py-1 rounded-md text-xs cursor-pointer">
            全部
          </div>
          <div className="bg-gray-100 px-2 py-1 rounded-md text-xs cursor-pointer">
            查看更多
          </div>
        </div>
      </div>
      {/* header end */}

      <div className="flex gap-6 my-4 min-h-0">
        <div className="w-1/2">
          <table className="w-full">
            <thead>
              <tr className="*:px-2 *:py-3 *:text-left *:font-normal *:text-[0.7rem] *:text-gray-400 *:uppercase">
                <th>序号</th>
                <th>创意名称</th>
                <th>发布日期</th>
                <th>售价</th>
              </tr>
            </thead>
            <tbody className="bg-white">
              {tableData.slice(0, 5).map((item) => (
                <tr className="*:p-2 *:whitespace-nowrap *:text-xs overflow-hidden cursor-pointer hover:bg-blue-50 rounded-md">
                  <td className="text-gray-400">{item.index}</td>
                  <td>{item.name}</td>
                  <td>{item.publishDate}</td>
                  <td>{item.sellPrice}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
        <div className="w-1/2">
          {tableData.length > 5 && (
            <table className="w-full">
              <thead>
                <tr className="*:px-2 *:py-3 *:text-left *:font-normal *:text-[0.7rem] *:text-gray-400 *:uppercase">
                  <th>序号</th>
                  <th>创意名称</th>
                  <th>发布日期</th>
                  <th>售价</th>
                </tr>
              </thead>
              <tbody>
                {tableData.slice(5).map((item) => (
                  <tr className="*:p-2 *:whitespace-nowrap *:text-xs  *:cursor-pointer rounded-md hover:bg-blue-50">
                    <td className="text-gray-400">{item.index}</td>
                    <td>{item.name}</td>
                    <td>{item.publishDate}</td>
                    <td>{item.sellPrice}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          )}
        </div>
      </div>
    </div>
  );
}
