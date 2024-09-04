"use client";
import { Table_Basic, TagType } from "@/app/types";
import { useState } from "react";
import { useRouter } from "next/navigation";
import ChangeFollowTopic from "./ChangeFollowTopic";
import Tab from "../Tab";
import { OrderTag } from "../CustomTag";

const topic: TagType[] = [
  { name: "推荐", key: 0 },
  { name: "电子信息", key: 1 },
  { name: "航空航天", key: 2 },
  { name: "人脸识别", key: 3 },
];

const types: TagType[] = [
  { name: "最新", key: 0 },
  { name: "最热", key: 1 },
  { name: "畅销", key: 2 },
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

export default function HomeSuggestion() {
  const router = useRouter();
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
  return (
    <div className="w-full my-8 p-10 bg-white rounded border border-primary-border">
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
                  className={`${currentTopic.key === item.key ? "text-white bg-primary" : "text-black bg-primary-light_bg/50"}`}
                  onClick={() => onSwitchTopic(item)}
                >
                  {item.name}
                </li>
              ))}
              <li
                className="text-black bg-primary-light_bg/50"
                onClick={() => setIsAdd(true)}
              >
                + 新增主题
              </li>
            </ul>
            <div className="flex gap-2">
              <Tab data={types} onChange={onChange} />
              <div className="bg-primary-light_bg/50 px-2 py-1 rounded-md text-[0.75rem] cursor-pointer">
                查看更多
              </div>
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
                  {tableData.slice(0, 5).map((item) => (
                    <tr
                      onClick={() => router.push(`/idea/${item.index}`)}
                      className="*:p-2 *:whitespace-nowrap *:text-[#333] overflow-hidden cursor-pointer hover:bg-blue-50 rounded-md"
                    >
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
                    {tableData.slice(5).map((item) => (
                      <tr
                        onClick={() => router.push(`/idea/${item.index}`)}
                        className="*:p-2 *:whitespace-nowrap *:text-base *:cursor-pointer *:text-[#333] rounded-md hover:bg-blue-50"
                      >
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
        </>
      )}
    </div>
  );
}
