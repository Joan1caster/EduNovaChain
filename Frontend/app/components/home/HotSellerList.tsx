"use client";
import { Table_Basic, TagType } from "@/app/types";
import { useState } from "react";
import Tab from "../Tab";
import { OrderTag } from "../CustomTag";

const types: TagType[] = [
  { name: "一天", id: 0 },
  { name: "一周", id: 1 },
  { name: "一月", id: 2 },
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
  const onChange = (item: TagType) => {
    //
  };
  return (
    <div className="w-full my-8 px-10 py-8 bg-white rounded border border-primary-border">
      {/* header start */}
      <div className="flex justify-between gap-2 items-center">
        <div className="text-lg">热门榜单</div>
        <div className="flex gap-2">
          {/* <Tab data={types} onChange={onChange} />
          <div className="bg-primary-light_bg px-2 py-1 rounded-md text-xs cursor-pointer">
            全部
          </div>
          <div className="bg-primary-light_bg px-2 py-1 rounded-md text-xs cursor-pointer">
            查看更多
          </div> */}
        </div>
      </div>
      {/* header end */}

      <div className="flex gap-6 my-4 min-h-0">
        <div className="w-1/2">
          <table className="w-full">
            <thead>
              <tr className="*:px-2 *:py-3 *:text-left *:font-normal *:text-sm *:text-primary-font_9 *:uppercase">
                <th>序号</th>
                <th>创意名称</th>
                <th>发布日期</th>
                <th>售价</th>
              </tr>
            </thead>
            <tbody className="bg-white">
              {tableData.slice(0, 5).map((item, i) => (
                <tr className="*:p-2 *:whitespace-nowrap overflow-hidden cursor-pointer hover:bg-blue-50 rounded-md">
                  <td>
                    <OrderTag order={item.index} bg={i < 3} />
                  </td>
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
                <tr className="*:px-2 *:py-3 *:text-left *:font-normal *:text-sm *:text-primary-font_9 *:uppercase">
                  <th>序号</th>
                  <th>创意名称</th>
                  <th>发布日期</th>
                  <th>售价</th>
                </tr>
              </thead>
              <tbody>
                {tableData.slice(5).map((item) => (
                  <tr className="*:p-2 *:whitespace-nowrap *:cursor-pointer rounded-md hover:bg-blue-50">
                    <td>
                      <OrderTag order={item.index} bg={false} />
                    </td>
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
