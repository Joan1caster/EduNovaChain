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
    <div>
      <div className="text-center my-6">New</div>
      <div className="relative w-full max-w-2xl mx-auto">
        <div className="overflow-hidden">
          <div
            className="flex justify-between transition-transform duration-500 ease-out"
            style={{ transform: `translateX(-${currentIndex * 33.3333}%)` }} // 修改为33.3333%来实现每次移动1/3
          >
            {slides.map((slide, index) => (
              <div className="flex-none w-1/3 pr-2" key={index}>
                <div className="rounded overflow-hidden">
                  <img
                    src={slide}
                    alt={`Slide ${index + 1}`}
                    className="w-full"
                  />
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
              className={`${currentIndex === i ? "w-4 bg-blue-400" : "w-1.5 bg-gray-400"} hover:bg-blue-200 h-1.5 rounded`}
            ></div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default Banner;
