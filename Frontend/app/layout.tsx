import type { Metadata } from "next";
import Image from "next/image";
import { Inter } from "next/font/google";
import "./globals.css";

import Link from "next/link";
import LoginButton from "./components/LoginButton";
import { cookies } from "next/headers";
import WagmiContext from "./context/WagmiContext";
import { ConnectButton } from "@rainbow-me/rainbowkit";
import Header from "./components/Header";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

export default function RootLayout({
  children,
  modal,
}: Readonly<{
  children: React.ReactNode;
  modal: React.ReactNode;
}>) {
  const email = cookies().get("email")?.value ?? "";

  return (
    <html lang="en">
      <body className={inter.className}>
        <WagmiContext>
          <Header />

          <section className="mx-auto max-w-4xl px-4 py-6 sm:px-6 sm:py-6 lg:max-w-7xl lg:px-8 text-xs">
            {/* <ConnectButton /> */}
            {children}
            {modal}
          </section>

          <footer className="p-4 w-full mt-6 bg-blue-700 text-gray-300">
            <div className="text-center text-sm font-light">
              Copyright © 2024 EduNovaChain.All Rights Reserved.
            </div>
          </footer>
        </WagmiContext>
      </body>
    </html>
  );
}
