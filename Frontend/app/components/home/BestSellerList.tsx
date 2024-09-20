"use client";

import { NFT, Table_Basic, TagType } from "@/app/types";
import { useState } from "react";
// import Tab from "../Tab";
import { AiTag, OrderTag } from "../CustomTag";
import { useAsyncEffect } from "ahooks";
import { useRouter } from "next/navigation";

const types: TagType[] = [
  { name: "一天", id: 0 },
  { name: "一周", id: 1 },
  { name: "一月", id: 2 },
];

export default function BestSellerList() {
  const router = useRouter();
  const [tableData, setTableData] = useState<NFT[]>([]);
  const onChange = (item: TagType) => {
    //
  };

  useAsyncEffect(async () => {
    try {
      const response = await (
        await fetch("/api/nfts?type=HighTrading&count=8")
      ).json();
      setTableData(response.data);
    } catch {
      //
    }
  }, []);
  return (
    <div className="w-full my-8 px-10 py-8 bg-white rounded border border-primary-border">
      {/* header start */}
      <div className="flex justify-between gap-2 items-center">
        <div className="text-lg">畅销榜单</div>
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

      <div className="grid grid-cols-4 gap-6 my-6">
        {tableData.map((item, i) => (
          <div
            key={i}
            className="p-6 border border-primary-border rounded-md"
            onClick={() => router.push(`/nft/${item.ID}`)}
          >
            <div className="flex gap-2 items-start">
              <div>
                <OrderTag order={i} />
              </div>
              <div className="flex-1">
                <div className="text-primary-font_3">
                  {item.Creator.Username}
                  {/* {item.isAi ? <AiTag /> : <></>} */}
                </div>
                <div className="my-6 flex items-center gap-2 text-sm font-light text-gray-400">
                  <p>售价：{item.Price}</p>
                  <p className="text-xs py-0.5 px-1 text-white bg-[#4BE2BB] rounded-sm">
                    {item.LikeCount}人支持
                  </p>
                </div>
                <div className="text-right text-sm font-light text-gray-300">
                  {item.CreatedAt}
                </div>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
