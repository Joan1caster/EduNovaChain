"use client";
import { Table_Basic, TagType } from "@/app/types";
import { useState } from "react";
import Tab from "../Tab";
import { AiTag, OrderTag } from "../CustomTag";

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
  const onChange = (item: TagType) => {
    //
  };
  return (
    <div className="w-full my-8 p-10 bg-white rounded border border-primary-border">
      {/* header start */}
      <div className="flex justify-between gap-2 items-center">
        <div className="text-lg">畅销榜单</div>
        <div className="flex gap-2">
          <Tab data={types} onChange={onChange} />
          <div className="bg-primary-light_bg px-2 py-1 rounded-md text-xs cursor-pointer">
            全部
          </div>
          <div className="bg-primary-light_bg px-2 py-1 rounded-md text-xs cursor-pointer">
            查看更多
          </div>
        </div>
      </div>
      {/* header end */}

      <div className="grid grid-cols-4 gap-6 my-6">
        {tableData.map((item, i) => (
          <div key={i} className="p-6 border border-primary-border rounded-md">
            <div className="flex gap-2 items-start">
              <div>
                <OrderTag order={item.index} />
              </div>
              <div className="flex-1">
                <div className="text-primary-font_3">
                  {item.name}
                  {item.isAi ? <AiTag /> : <></>}
                </div>
                <div className="my-6 flex items-center gap-2 text-sm font-light text-gray-400">
                  <p>售价：{item.sellPrice}</p>
                  <p className="text-xs py-0.5 px-1 text-white bg-[#4BE2BB] rounded-sm">
                    199人支持
                  </p>
                </div>
                <div className="text-right text-sm font-light text-gray-300">
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
