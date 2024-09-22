"use client";

import { NFT } from "@/app/types";
import { useAsyncEffect } from "ahooks";
import Link from "next/link";
import { useState, useEffect } from "react";
import { BannerCard } from "../CustomTag";

let interval: NodeJS.Timeout;

const Banner = () => {
  const [currentIndex, setCurrentIndex] = useState(0);
  const [total, setTotal] = useState(0);
  const [data, setData] = useState<NFT[]>([]);

  useEffect(() => {
    clearInterval(interval);
    interval = setInterval(() => {
      setCurrentIndex((prevIndex) =>
        prevIndex + 1 >= total ? 0 : prevIndex + 1
      );
    }, 3000);

    return () => clearInterval(interval);
  }, [total, currentIndex]);

  useAsyncEffect(async () => {
    try {
      const response = await (await fetch("/api/nfts?type=latest&count=6")).json();
      setTotal(response.count);
      setData(response.data);
    } catch {
      //
    }
  }, []);

  const showSlide = (index: number) => {
    setCurrentIndex(index);
  };

  return (
    <div className="w-full">
      <div className="w-[332px] h-[62px] mx-auto bg-[url('/images/slice/banner_title.png')] bg-no-repeat bg-cover"></div>
      <div className="relative w-full h-72 mt-9">
        <div className="overflow-hidden">
          <div
            className="flex gap-x-6 flex-nowrap w-full transition-transform duration-500 ease-out"
            style={{ transform: `translateX(-${currentIndex * 524}px)` }} // 修改为33.3333%来实现每次移动1/3
          >
            {data.map((item: NFT, index: number) => (
              <BannerCard order={index} key={index}>
                  <Link href={`/nft/${item.ID}`}>
                    <p className="text-[20px] text-[#293748]">{item.Title}</p>
                    <p className="my-6 text-[14px] text-[#666]">
                      {item.Price}ETH
                    </p>
                    <p className="absolute right-6 bottom-6 text-[14px] text-right text-[#ABC5EB] font-light">
                      {item.CreatedAt}
                    </p>
                  </Link>
              </BannerCard>
            ))}
          </div>
        </div>

        <div className="flex justify-center mt-2 space-x-2">
          {new Array(total).fill(1).map((_, i) => (
            <div
              key={i}
              onClick={() => showSlide(i)}
              className={`${currentIndex === i ? "w-6 bg-primary" : "w-2 bg-primary/20"} hover:bg-blue-200 h-2 rounded cursor-pointer`}
            ></div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default Banner;
