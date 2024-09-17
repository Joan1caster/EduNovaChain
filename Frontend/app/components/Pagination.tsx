"use client";
import { useEffect, useState } from "react";

interface PaginationProps {
  total: number;
  currentPage: number;
  onPageChange: (page: number) => void;
}

const Pagination: React.FC<PaginationProps> = ({
  total,
  currentPage,
  onPageChange,
}) => {
  const [current, setCurrent] = useState(currentPage);
  const [totalPage, setTotalPage] = useState<number>(0);
  const [pages, setPages] = useState<(number | string)[]>([]);

  const handlePageClick = (page: number) => {
    setCurrent(page);
    onPageChange(page);
  };

  const onClickQuote = (i: number) => {
    const newCurrent = Math.ceil(
      (pages[i - 1] as number) + (pages[i + 1] as number) / 2
    );
    setCurrent(newCurrent);
    onPageChange(newCurrent);
  };

  const generatePages = () => {
    const newPages: (number | string)[] = [];
    const pageLimit = 10;

    if (totalPage <= pageLimit) {
      for (let i = 1; i <= totalPage; i++) {
        newPages.push(i);
      }
    } else {
      newPages.push(1);

      if (current <= 4) {
        for (let i = 2; i <= 5; i++) {
          newPages.push(i);
        }
        newPages.push("...");
        newPages.push(totalPage);
      } else if (current >= totalPage - 3) {
        newPages.push("...");
        for (let i = totalPage - 4; i < totalPage; i++) {
          newPages.push(i);
        }
        newPages.push(totalPage);
      } else {
        newPages.push("...");
        for (let i = current - 1; i <= current + 1; i++) {
          newPages.push(i);
        }
        newPages.push("...");
        newPages.push(totalPage);
      }
    }
    setPages(newPages);
  };

  useEffect(() => {
    generatePages();
  }, [current, totalPage]);

  useEffect(() => {
    setTotalPage(Math.ceil(total / 20));
  }, [total]);
  return (
    <div className="w-full absolute bottom-4">
      <div className="w-80 mx-auto text-center">
        {pages.map((item, i) =>
          typeof item === "number" ? (
            <p
              key={`${item}_${i}`}
              className={`w-10 h-10 inline-block leading-10 rounded ${current === item ? " bg-primary text-white" : "text-[#888D95]"} cursor-pointer`}
              onClick={() => handlePageClick(item)}
            >
              {item}
            </p>
          ) : (
            <p
              key={`${item}_${i}`}
              className={`w-10 h-10 inline-block leading-10 rounded text-[#888D95] cursor-pointer`}
              onClick={() => onClickQuote(i)}
            >
              {item}
            </p>
          )
        )}
      </div>
    </div>
  );
};

export default Pagination;
