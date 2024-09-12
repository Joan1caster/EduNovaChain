"use client";

import Link from "next/link";
import Image from "next/image";
import { usePathname } from "next/navigation";

export default function Nav() {
  const pathname = usePathname();

  return (
    <div className="flex-1 ml-6 flex gap-6 items-center">
      <Link
        href="/"
        className={`text-md ${pathname === "/" ? "text-black" : "text-primary-font_9"} hover:text-black`}
      >
        主页
      </Link>
      <Link
        href="/category"
        className={`text-md ${pathname === "/category" ? "text-black" : "text-primary-font_9"} hover:text-black`}
      >
        分类
      </Link>

      <div className="relative group/item text-md text-primary-font_9 cursor-pointer hover:text-black">
        发布
        <Image
          src={"/images/slice/AI.png"}
          alt={"AI.png"}
          width={20}
          height={42}
          className="absolute -top-1 left-9"
        />
        <div className="hidden static z-10 w-28 text-center group-hover/item:block group-hover/item:absolute bg-white shadow-lg rounded-md p-1">
          <Link
            className="block rounded-lg py-2 px-3 transition hover:bg-gray-100"
            href="/create"
          >
            <p className="text-primary-font_9 text-xs">发布创意</p>
          </Link>
          <Link
            className="block rounded-lg py-2 px-3 transition hover:bg-gray-100"
            href="/create"
          >
            <p className="text-primary-font_9 text-xs">AI生成创意</p>
          </Link>
        </div>
      </div>
    </div>
  );
}
