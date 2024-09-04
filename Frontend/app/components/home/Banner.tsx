"use client";
import { useState, useEffect } from "react";
let interval: NodeJS.Timeout;
const Banner = () => {
  const [currentIndex, setCurrentIndex] = useState(0);
  const totalSlides = 3; // 幻灯片数量

  const slides = [
    "https://via.placeholder.com/800x400?text=Slide+1",
    "https://via.placeholder.com/800x400?text=Slide+2",
    "https://via.placeholder.com/800x400?text=Slide+3",
    "https://via.placeholder.com/800x400?text=Slide+1",
    "https://via.placeholder.com/800x400?text=Slide+2",
    "https://via.placeholder.com/800x400?text=Slide+3",
  ];

  useEffect(() => {
    clearInterval(interval);
    interval = setInterval(() => {
      setCurrentIndex((prevIndex) => (prevIndex + 1) % totalSlides);
    }, 3000); // 每3秒切换一次

    return () => clearInterval(interval); // 清除计时器
  }, [totalSlides, currentIndex]);

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
            className="flex justify-between transition-transform duration-500 ease-out"
            style={{ transform: `translateX(-${currentIndex * 33.3333}%)` }} // 修改为33.3333%来实现每次移动1/3
          >
            {slides.map((slide, index) => (
              <div className="flex-none w-1/3 h-[160px] pr-3" key={index}>
                <div className="relative h-full p-6 rounded overflow-hidden bg-[url('/images/slice/banner_card_bg.jpg')] bg-cover bg-no-repeat">
                  <p className="text-[20px] text-[#293748]">创意点子标题内容</p>
                  <p className="my-6 text-[14px] text-[#666]">免费</p>
                  <p className="absolute right-6 bottom-6 text-[14px] text-right text-[#ABC5EB] font-light">
                    2024-09-04
                  </p>
                </div>
              </div>
            ))}
          </div>
        </div>

        <div className="flex justify-center mt-2 space-x-2">
          {new Array(totalSlides).fill(1).map((_, i) => (
            <div
              onClick={() =>
                showSlide((currentIndex - 1 + totalSlides) % totalSlides)
              }
              className={`${currentIndex === i ? "w-6 bg-primary" : "w-2 bg-primary/20"} hover:bg-blue-200 h-2 rounded`}
            ></div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default Banner;
