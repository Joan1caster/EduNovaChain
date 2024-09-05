"use client";

import { usePathname } from "next/navigation";
import { useEffect, useState } from "react";

export default function Footer() {
  const pathname = usePathname();

  const [visible, setVisible] = useState(false);

  useEffect(() => {
    setVisible(!/^\/account\/./.test(pathname));
  }, [pathname]);

  if (visible)
    return (
      <footer className="p-4 w-full mt-6 bg-primary text-gray-300">
        <div className="text-center text-sm font-light">
          Copyright © 2024 {process.env.NEXT_PUBLIC_APP_NAME}.All Rights
          Reserved.
        </div>
      </footer>
    );
  else return <></>;
}
