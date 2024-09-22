"use client";
import Image from "next/image";
import { IdeaInfo, NFT, Table_Basic, TagType } from "@/app/types";
import { useEffect, useState } from "react";
import { useWriteContract } from "wagmi";
import { ContractConfig } from "@/app/abis";
import { parseEther } from "viem";

const info_tag: TagType[] = [
  { name: "基本信息", id: 0 },
  { name: "创意详设", id: 1 },
];
const types: TagType[] = [
  { name: "引用记录", id: 0 },
  { name: "被引用记录", id: 1 },
];

export default function Page({ params: { id } }: { params: { id: string } }) {
  const [nftInfo, setNFTInfo] = useState<NFT>();
  const [currentTag, setCurrentTag] = useState<TagType>(info_tag[0]);
  const [currentType, setCurrentType] = useState<TagType>(types[0]);
  const { isError, isPending, writeContractAsync } = useWriteContract();
  const [loading, setLoading] = useState(false);

  const onSwitchTag = (item: TagType) => {
    setCurrentTag(item);
  };
  const onSwitchType = (item: TagType) => {
    setCurrentType(item);
  };

  const getDetail = async () => {
    const response = await fetch(`/api/nft?id=${id}`);
    const data = (await response.json()) as NFT;
    setNFTInfo(data);
  };

  const onBuy = async () => {
    setLoading(true);
    try {
      const orderResponse = await fetch(`/api/order?type=orderId&value=${id}`);
      const orderJson = await orderResponse.json();
      const orderId = orderJson.data.ID;

      console.log(parseEther(orderJson.data.Price.toString()));
      const hash = await writeContractAsync({
        ...ContractConfig,
        functionName: "purchaseInnovation",
        args: [nftInfo?.TokenID],
        value: parseEther(orderJson.data.Price.toString()),
      });
      console.log(hash, parseEther(orderJson.data.Price.toString()));

      // const hash =
      //   "0xde64c521638a1f32724abd6e0040e857197196c314569a16af04b627d963ea32";

      const buyResponse = await fetch("/api/order", {
        method: "POST",
        body: JSON.stringify({
          order_id: orderId,
          tx_hash: hash,
        }),
      });
      console.log(buyResponse);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    getDetail();
  }, []);

  return (
    <div>
      {/* head start */}
      <h3 className="text-lg mt-2">{nftInfo?.Title}</h3>
      <div className="flex justify-between items-center my-2">
        <div className="flex items-center gap-1">
          <Image
            src="/images/png/cute_devil.png"
            width={12}
            height={12}
            alt={nftInfo?.Creator.Username || ""}
          />
          <p className="text-sm text-blue-400">{nftInfo?.Creator.Username}</p>
        </div>
        <div className="flex items-end">
          <span className="text-base leading-none">{nftInfo?.LikeCount}</span>
          <span className="text-[0.6rem] text-gray-400 ml-1 leading-none">
            获赞
          </span>
          <span className="text-base leading-none ml-4">
            {nftInfo?.ViewCount}
          </span>
          <span className="text-[0.6rem] text-gray-400 ml-1 leading-none">
            浏览
          </span>
          <span className="text-base leading-none ml-4">
            {nftInfo?.TransactionCount}
          </span>
          <span className="text-[0.6rem] text-gray-400 ml-1 leading-none">
            购买
          </span>
          {/* <span className="text-base leading-none ml-4">{info.citation}</span>
          <span className="text-[0.6rem] text-gray-400 ml-1 leading-none">
            被引用
          </span> */}
        </div>
      </div>
      {/* head end */}

      {/* basic info start */}
      <div className="h-80 flex gap-4 my-4">
        {/* left start */}
        <div className="h-full flex-1 bg-white rounded-sm shadow-sm p-6">
          <div className="flex justify-between">
            <ul className="flex w-36 justify-around bg-gray-100 rounded-md *:text-sm *:hover:cursor-pointer">
              {info_tag.map((item) => (
                <li
                  className={`w-full py-1 rounded-md text-center ${currentTag.id === item.id ? " bg-gray-200" : " bg-gray-100"}`}
                  onClick={() => onSwitchTag(item)}
                >
                  {item.name}
                </li>
              ))}
            </ul>
            <div className="flex items-center">
              {/* <svg
                t="1725261948952"
                className="icon"
                viewBox="0 0 1024 1024"
                version="1.1"
                xmlns="http://www.w3.org/2000/svg"
                p-id="1225"
                width="14"
                height="14"
              >
                <path
                  d="M448 864C218.6 864 32 677.4 32 448S218.6 32 448 32s416 186.6 416 416-186.6 416-416 416z m0-768C253.9 96 96 253.9 96 448s157.9 352 352 352 352-157.9 352-352S642.1 96 448 96z"
                  fill="#92a3af"
                  p-id="1226"
                ></path>
                <path
                  d="M960 992c-8.2 0-16.4-3.1-22.6-9.4l-224-224c-12.5-12.5-12.5-32.8 0-45.3s32.8-12.5 45.3 0l224 224c12.5 12.5 12.5 32.8 0 45.3-6.3 6.3-14.5 9.4-22.7 9.4z"
                  fill="#92a3af"
                  p-id="1227"
                ></path>
              </svg>
              <p className="text-[0.6rem] text-gray-400 ml-2">找相似</p> */}
            </div>
          </div>

          {currentTag.id === 0 ? (
            <div className="my-4">
              {/* basic info start */}
              <h4 className="mt-4">发布日期</h4>
              <p className="mt-1 text-sm text-gray-600">{nftInfo?.CreatedAt}</p>

              <h4 className="mt-4">学科分类</h4>
              <p className="mt-1 text-sm text-gray-600">
                {nftInfo?.Categories}
              </p>

              <h4 className="mt-4">摘要内容</h4>
              <p className="mt-1 text-sm text-gray-600">{nftInfo?.Summary}</p>

              <h4 className="mt-4">主题关键词</h4>
              <p className="mt-1 text-sm text-gray-600">
                {nftInfo?.Topics.map((item) => (
                  <span className="text-blue-400 after:content-[';'] last:after:content-none">
                    {item.Name}
                  </span>
                ))}
              </p>
              {/* basic info end */}
            </div>
          ) : (
            <div className="my-4">
              {/* idea info start */}
              {nftInfo?.Content}
              {/* idea info end */}
            </div>
          )}
        </div>
        {/* left end */}
        {/* right start */}
        <div className="h-full w-60 bg-white rounded-sm shadow-sm p-4">
          <h1 className="text-center font-bold text-black my-4">最低售价</h1>
          <p className="text-center text-[0.6rem] text-gray-200">
            当前作品为定价销售，可直接按照以下价格一次性购买使用
          </p>
          <h1 className="text-center font-bold text-lg mt-6">
            {nftInfo?.Price}ETH
          </h1>
          {/* <p className="text-center text-base text-gray-200 my-2">
            ¥ {info.rmb}
          </p> */}
          <div className="mt-28">
            <button
              disabled={loading}
              onClick={onBuy}
              className="w-full rounded-md bg-blue-600 px-6 py-1.5 text-sm font-light text-white shadow-sm hover:bg-blue-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-blue-600 disabled:bg-blue-400 disabled:cursor-not-allowed"
            >
              购买
            </button>
            {/* <button className="w-full rounded-md bg-gray-100 my-2 px-6 py-1.5 text-sm font-light text-gray-600 shadow-sm hover:bg-gray-200 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-gray-200">
              加入购物车
            </button> */}
          </div>
        </div>
        {/* right end */}
      </div>
      {/* basic info end */}

      {/* record start */}
      <div className="min-h-52 my-4">
        <div className=" bg-white rounded-sm shadow-sm p-4">
          <div className="w-full">
            <h1 className="text-md mb-2">交易记录</h1>
            <table className="w-full min-h-52">
              <thead>
                <tr className="*:px-2 *:py-3 *:text-left *:font-normal *:text-sm *:text-gray-400 *:uppercase">
                  <th>序号</th>
                  <th>交易价格</th>
                  <th>人民币价格</th>
                  <th>交易日期</th>
                  <th>交易者</th>
                </tr>
              </thead>
              <tbody className="bg-white">
                {(nftInfo?.HistoryList ?? []).slice(0, 5).map((item) => (
                  <tr className="*:p-2 *:whitespace-nowrap *:text-xs overflow-hidden cursor-pointer hover:bg-blue-50 rounded-md">
                    <td className="text-gray-400">{item.index}</td>
                    <td>{item.name}</td>
                    <td>{item.publishDate}</td>
                    <td>{item.sellPrice}</td>
                    <td>{item.sellPrice}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
        {/* <div className=" bg-white rounded-sm shadow-sm p-4">
          <div className="w-full">
            <div className="flex justify-between items-end mb-2">
              <h1 className="text-md">{currentType.name}</h1>
              <ul className="flex w-36 justify-around bg-gray-100 rounded-md *:text-sm *:hover:cursor-pointer">
                {types.map((item) => (
                  <li
                    className={`w-full py-1 rounded-md text-center ${currentType.key === item.key ? " bg-gray-200" : " bg-gray-100"}`}
                    onClick={() => onSwitchType(item)}
                  >
                    {item.name}
                  </li>
                ))}
              </ul>
            </div>
            <table className="w-full">
              <thead>
                <tr className="*:px-2 *:py-3 *:text-left *:font-normal *:text-sm *:text-gray-400 *:uppercase">
                  <th>序号</th>
                  <th>引用创意名称</th>
                  <th>作者</th>
                  <th>售价</th>
                </tr>
              </thead>
              <tbody className="bg-white">
                {tableData.slice(0, 3).map((item) => (
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
        </div> */}
      </div>
      {/* record end */}

      {/* interaction start */}
      <div className="flex justify-center gap-6 relative my-6">
        <div
          className={`w-12 h-12 rounded-full flex items-center flex-col justify-center border border-gray-100 cursor-pointer ${nftInfo?.Likes ? "bg-blue-400 text-white" : "focus:bg-blue-200"}`}
        >
          <svg
            t="1725256924681"
            className="icon"
            viewBox="0 0 1024 1024"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            p-id="1961"
            width="16"
            height="16"
          >
            <path
              d="M489.264 908.336L157.7728 576.8416c-84.352-84.352-84.352-221.1168 0-305.4704 84.352-84.352 221.1168-84.352 305.4704 0l48.7552 48.7552 48.7584-48.7552c83.5104-83.5088 218.3872-84.344 302.9248-2.5056l2.544 2.5056c84.3552 84.352 84.3552 221.1168 0 305.4704L534.736 908.3344a31.9152 31.9152 0 0 1-22.9104 9.3744h-0.216a31.9104 31.9104 0 0 1-22.3456-9.3744zM512 863.1856l320.2864-320.2848c65.5648-65.5648 65.608-171.84 0.128-237.4592l-2.2512-2.216c-65.824-63.5872-170.6336-62.744-235.4656 2.088l-48.7584 48.7552c-18.744 18.744-49.136 18.744-67.8816 0l-48.7552-48.7568c-65.608-65.608-171.9808-65.608-237.5888 0-65.608 65.608-65.608 171.9808 0 237.5888L512 863.1872z"
              fill={`${nftInfo?.Likes ? "#fff" : "#222"}`}
              p-id="1962"
            ></path>
          </svg>
          <p
            className={`text-[0.6rem] ${nftInfo?.Likes ? "text-white" : "text-gray-400"}`}
          >
            {nftInfo?.LikeCount}
          </p>
        </div>
        {/* <div
          className={`w-12 h-12 rounded-full flex items-center flex-col justify-center border border-gray-100 cursor-pointer ${info.userCitationed ? "bg-blue-400 text-white" : "focus:bg-blue-200"}`}
        >
          <svg
            t="1725256872332"
            className="icon"
            viewBox="0 0 1024 1024"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            p-id="3313"
            width="16"
            height="16"
          >
            <path
              d="M543.5952 384c13.2544 0 24 10.7456 24 24 0 13.2544-10.7456 24-24 24H424c-8.8368 0-16 7.1632-16 16v384c0 8.8368 7.1632 16 16 16h472c8.8368 0 16-7.1632 16-16V448c0-8.8368-7.1632-16-16-16H735.664c-13.2528 0-24-10.7456-24-24 0-13.2544 10.7472-24 24-24H896c35.3456 0 64 28.6544 64 64v384c0 35.3456-28.6544 64-64 64H424c-35.3456 0-64-28.6544-64-64V448c0-35.3456 28.6544-64 64-64zM600 128c35.3456 0 64 28.6544 64 64v384c0 35.3456-28.6544 64-64 64h-119.5952c-13.2544 0-24-10.7456-24-24 0-13.2544 10.7456-24 24-24H600c8.8368 0 16-7.1632 16-16V192c0-8.8368-7.1632-16-16-16H128c-8.8368 0-16 7.1632-16 16v384c0 8.8368 7.1632 16 16 16h160.336c13.2528 0 24 10.7456 24 24 0 13.2544-10.7472 24-24 24H128c-35.3456 0-64-28.6544-64-64V192c0-35.3456 28.6544-64 64-64z"
              fill={`${info.userCitationed ? "#fff" : "#222"}`}
              p-id="3314"
            ></path>
          </svg>
          <p
            className={`text-[0.6rem] ${info.userCitationed ? "text-white" : "text-gray-400"}`}
          >
            引用
          </p>
        </div> */}
        <div className="absolute right-0 -bottom-2 flex">
          <svg
            t="1725257786425"
            className="icon"
            viewBox="0 0 1024 1024"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            p-id="2110"
            width="16"
            height="16"
          >
            <path
              d="M544.7296 112.5856a64 64 0 0 1 23.4256 23.424l383.1296 663.6c17.6736 30.6112 7.184 69.7536-23.4256 87.4272a64 64 0 0 1-32 8.5744H129.6c-35.3456 0-64-28.6544-64-64a64 64 0 0 1 8.5744-32l383.1296-663.6c17.6736-30.6112 56.816-41.0992 87.4256-23.4256z m-45.856 47.424L115.744 823.6112a16 16 0 0 0-2.144 8c0 8.8384 7.1632 16 16 16h766.2592a16 16 0 0 0 8-2.144c7.6528-4.416 10.2752-14.2016 5.856-21.856l-383.1296-663.6a16 16 0 0 0-5.856-5.856c-7.6528-4.4176-17.4384-1.7952-21.856 5.856zM512 672c26.5104 0 48 21.4896 48 48 0 26.5104-21.4896 48-48 48-26.5104 0-48-21.4896-48-48 0-26.5104 21.4896-48 48-48z m0-352c13.2544 0 24 10.7456 24 24v256c0 13.2544-10.7456 24-24 24-13.2544 0-24-10.7456-24-24V344c0-13.2544 10.7456-24 24-24z"
              fill="#92a3af"
              p-id="2111"
            ></path>
          </svg>
          <p className="text-[0.6rem] text-gray-400 ml-1">举报</p>
        </div>
      </div>
      {/* interaction end */}

      {/* comment start */}
      <form className="w-full my-4">
        <textarea
          className="w-full border-gray-100 rounded-sm"
          rows={3}
          maxLength={300}
        />
        <div className="flex justify-end mt-4">
          <button
            type="submit"
            disabled={true}
            className="rounded-md bg-blue-600 px-6 py-1 text-sm font-light text-white shadow-sm hover:bg-blue-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-blue-600 disabled:bg-blue-400 disabled:cursor-not-allowed"
          >
            发表
          </button>
        </div>
      </form>
      {/* comment end */}

      <div className="my-4">
        <div className="flex items-start gap-2">
          <Image
            src="/images/png/devil_laugh.png"
            alt="avator"
            width={36}
            height={36}
            className="rounded-full border border-gray-100"
          />
          <div className="flex-1">
            <div>
              <p>
                莉丝{" "}
                <span className="text-[0.6rem] text-gray-300 font-light inline-block pl-2">
                  2024-09-02
                </span>
              </p>
            </div>
            <div className="my-2 text-sm text-gray-400">评论</div>
            <div className="flex justify-end items-center">
              <svg
                t="1725259764914"
                className="icon"
                viewBox="0 0 1024 1024"
                version="1.1"
                xmlns="http://www.w3.org/2000/svg"
                p-id="4394"
                width="16"
                height="16"
              >
                <path
                  d="M512.6 992c-18 0-36.2-1-54.5-3C254.8 966.9 61.4 776.9 35.9 574.2c-18.8-149.5 31-295.9 136.7-401.6C278.2 66.9 424.7 17 574.2 35.9 776.6 61.4 966.5 254.5 989 457.5c12.1 109.5-11.8 215.6-69.1 307.7 12.1 44.1 46 161 46 161 3.2 11.2 0.1 23.3-8.1 31.5s-20.3 11.4-31.5 8.1c0 0-117.2-33.9-160.5-46.3-76.7 47.9-163.2 72.5-253.2 72.5z m-0.3-896c-110.5 0-215.6 43-294.5 121.8C126.2 309.4 83 436.4 99.4 566.2 121.1 738.7 292 906.6 465.1 925.4c97.8 10.7 192.3-11.7 273.5-64.5 10.6-6.9 23.9-8.8 36.4-5.3 19.3 5.5 71.9 20.7 113.3 32.7-12-41.7-27.2-94.9-32.6-115-3.2-12.1-1.1-25 5.8-35.5 52.5-81 74.6-175.5 63.8-273.2-19-172.9-186.9-343.5-359.1-365.2-18-2.3-36.1-3.4-53.9-3.4z m261.2 818.5z m141.7-142.1c0 0.1-0.1 0.1-0.1 0.2 0.1 0 0.1-0.1 0.1-0.2z"
                  fill="#92a3af"
                  p-id="4395"
                ></path>
                <path
                  d="M336 544c-26.5 0-48-21.5-48-48s21.5-48 48-48 48 21.5 48 48-21.6 48-48 48zM528 544c-26.5 0-48-21.5-48-48s21.5-48 48-48 48 21.5 48 48-21.6 48-48 48zM720 544c-26.5 0-48-21.5-48-48s21.5-48 48-48 48 21.5 48 48-21.6 48-48 48z"
                  fill="#92a3af"
                  p-id="4396"
                ></path>
              </svg>

              <svg
                t="1725256924681"
                className="icon ml-2"
                viewBox="0 0 1024 1024"
                version="1.1"
                xmlns="http://www.w3.org/2000/svg"
                p-id="1961"
                width="18"
                height="18"
              >
                <path
                  d="M489.264 908.336L157.7728 576.8416c-84.352-84.352-84.352-221.1168 0-305.4704 84.352-84.352 221.1168-84.352 305.4704 0l48.7552 48.7552 48.7584-48.7552c83.5104-83.5088 218.3872-84.344 302.9248-2.5056l2.544 2.5056c84.3552 84.352 84.3552 221.1168 0 305.4704L534.736 908.3344a31.9152 31.9152 0 0 1-22.9104 9.3744h-0.216a31.9104 31.9104 0 0 1-22.3456-9.3744zM512 863.1856l320.2864-320.2848c65.5648-65.5648 65.608-171.84 0.128-237.4592l-2.2512-2.216c-65.824-63.5872-170.6336-62.744-235.4656 2.088l-48.7584 48.7552c-18.744 18.744-49.136 18.744-67.8816 0l-48.7552-48.7568c-65.608-65.608-171.9808-65.608-237.5888 0-65.608 65.608-65.608 171.9808 0 237.5888L512 863.1872z"
                  fill="#92a3af"
                  p-id="1962"
                ></path>
              </svg>
              <p className="text-[0.6rem] text-gray-400">{nftInfo?.Likes}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
