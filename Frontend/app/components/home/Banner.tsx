"use client";

import { NFT } from "@/app/types";
import { useAsyncEffect } from "ahooks";
import Link from "next/link";
import { useState, useEffect } from "react";

let interval: NodeJS.Timeout;

const Banner = () => {
  const [currentIndex, setCurrentIndex] = useState(0);
  const [total, setTotal] = useState(0);
  const [data, setData] = useState<NFT[]>([]);
  const totalSlides = 3; // 幻灯片数量

  useEffect(() => {
    clearInterval(interval);
    interval = setInterval(() => {
      setCurrentIndex((prevIndex) => (prevIndex + 1) % total);
    }, 3000); // 每3秒切换一次

    return () => clearInterval(interval); // 清除计时器
  }, [totalSlides, currentIndex]);

  useAsyncEffect(async () => {
    try {
      const response = await (await fetch("/api/nfts")).json();
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
    <div className="w-full bg-[url('/images/slice/banner_bg.png')]">
      <div className="banner_title">
        <div className="new_EN">New</div>
        <div className="new_ZH">最新</div>
      </div>
      <div className="relative w-full">
        <div className="overflow-hidden">
          <div
            className="flex transition-transform duration-500 ease-out"
            style={{ transform: `translateX(-${currentIndex * 33.3333}%)` }} // 修改为33.3333%来实现每次移动1/3
          >
            {data.map((slide: NFT, index: number) => (
              <div className="flex-none w-1/3 h-[160px] pr-3" key={index}>
                <div className="relative h-full p-6 rounded overflow-hidden bg-[url('/images/slice/banner_card_bg.jpg')] bg-cover bg-no-repeat">
                  <Link href={`/nft/${slide.ID}`}>
                    <p className="text-[20px] text-[#293748]">
                      创意点子标题内容
                    </p>
                    <p className="my-6 text-[14px] text-[#666]">免费</p>
                    <p className="absolute right-6 bottom-6 text-[14px] text-right text-[#ABC5EB] font-light">
                      {slide.CreatedAt}
                    </p>
                  </Link>
                </div>
              </div>
            ))}
          </div>
        </div>

        <div className="flex justify-center mt-2 space-x-2">
          {new Array(total).fill(1).map((_, i) => (
            <div
              key={i}
              onClick={() => showSlide((currentIndex - 1 + total) % total)}
              className={`${currentIndex === i ? "w-6 bg-primary" : "w-2 bg-primary/20"} hover:bg-blue-200 h-2 rounded cursor-pointer`}
            ></div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default Banner;
