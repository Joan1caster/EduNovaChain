"use client";
import { TagType } from "@/app/types";
import { useState } from "react";
import Tab from "../Tab";
import { useAsyncEffect } from "ahooks";
import { TopicCard } from "../CustomTag";

const topic: TagType[] = [
  { name: "数学", id: 0 },
  { name: "物理与实验", id: 1 },
  { name: "生物科学", id: 2 },
  { name: "有机化学", id: 3 },
  { name: "写作", id: 4 },
];
const types: TagType[] = [
  { name: "小学", id: 0 },
  { name: "初中", id: 1 },
  { name: "高中", id: 2 },
  { name: "高校", id: 3 },
];
export default function Topic() {
  const [gradeList, setGradeList] = useState<TagType[]>([]);
  const [subjectList, setSubjectList] = useState<TagType[]>([]);

  useAsyncEffect(async () => {
    const response = await (await fetch("/api/grade")).json();
    if (response.count > 0) {
      setGradeList(response.data);
      onChange(response.data[0]);
    }
  }, []);

  const onChange = async (grade: TagType) => {
    const response = await (await fetch(`/api/subject?id=${grade.id}`)).json();
    if (response.count > 0) {
      setSubjectList(response.data);
    }
  };
  return (
    <div className="w-full my-8 px-10 py-8 bg-white rounded border border-primary-border">
      {/* header start */}
      <div className="flex justify-between gap-2 items-center">
        <div className="text-lg">专题广场</div>
        <div className="flex gap-2">
          <Tab data={gradeList} onChange={onChange} />
          {/* <div className="bg-primary-light_bg px-2 py-1 rounded-md text-xs cursor-pointer">
            查看更多
          </div> */}
        </div>
      </div>
      {/* header end */}

      <div className="grid grid-cols-5 gap-4 h-[240px] mt-8 mb-2">
        {subjectList.map((item, i) => (
          <TopicCard order={i} key={i}>
            <div className="absolute bottom-0 w-full py-4 text-center text-[#333] font-light bg-white/30">
              {item.name}
            </div>
          </TopicCard>
        ))}
      </div>
    </div>
  );
}
