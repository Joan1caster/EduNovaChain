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
    isAi: true,
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
    isAi: true,
  },
];

export default function BestSellerList() {
  const [currentType, setCurrentType] = useState<TagType>(types[0]);

  const onSwitchType = (item: TagType) => {
    setCurrentType(item);
  };
  return (
    <div className="w-full my-4 p-4 bg-white rounded shadow-sm">
      {/* header start */}
      <div className="flex justify-between gap-2 items-center">
        <div>畅销榜单</div>
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

      <div className="grid grid-cols-4 gap-4 my-4">
        {tableData.map((item, i) => (
          <div key={i} className="p-4 border border-gray-100 rounded-md">
            <div className="flex gap-2 items-start">
              <div>{item.index}</div>
              <div className="flex-1">
                <p>{item.name}</p>
                <div className="my-3 flex items-center gap-2 text-[0.7rem] font-light text-gray-400">
                  <p>售价：{item.sellPrice}</p>
                  <p className="text-[0.6rem] py-0.5 px-1 text-white bg-green-700 rounded-sm">
                    199人支持
                  </p>
                </div>
                <div className="text-right text-[0.6rem] font-light text-gray-300">
                  {item.publishDate}
                </div>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
