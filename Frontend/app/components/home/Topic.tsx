"use client";
import { TagType } from "@/app/types";
import { useState } from "react";
import Tab from "../Tab";

const topic: TagType[] = [
  { name: "数学", key: 0 },
  { name: "物理与实验", key: 1 },
  { name: "生物科学", key: 2 },
  { name: "有机化学", key: 3 },
  { name: "写作", key: 4 },
];
const types: TagType[] = [
  { name: "小学", key: 0 },
  { name: "初中", key: 1 },
  { name: "高中", key: 2 },
  { name: "高校", key: 3 },
];
export default function Topic() {
  const onChange = (item: TagType) => {
    //
  };
  return (
    <div className="w-full my-8 p-10 bg-white rounded border border-primary-border">
      {/* header start */}
      <div className="flex justify-between gap-2 items-center">
        <div className="text-lg">专题广场</div>
        <div className="flex gap-2">
          <Tab data={types} onChange={onChange} />
          <div className="bg-primary-light_bg px-2 py-1 rounded-md text-xs cursor-pointer">
            查看更多
          </div>
        </div>
      </div>
      {/* header end */}

      <div className="grid grid-cols-5 gap-4 h-32 my-4">
        {topic.map((item) => (
          <div className="relative bg-blue-200 rounded">
            <div className="absolute bottom-0 w-full py-2 text-center text-gray-700 font-light bg-white/20">
              {item.name}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
