import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
    "./node_modules/@vercel/examples-ui/**/*.js",
  ],
  theme: {
    extend: {
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
      },
      colors: {
        primary: {
          light: "",
          DEFAULT: "#1474FC",
          dark: "",
          light_bg: "#E9EFF6",
          border: "#F0F6FF",
          font_3: "#333",
          font_9: "#999",
          font_C: "#CCC",
        },
      },

      keyframes: {
        // 自定义动画: 圆角柱从长到短，透明度从小到大
        loading: {
          "0%, 100%": { height: "10px", opacity: "0.1" },
          "25%": { height: "10px", opacity: "0.2" },
          "50%": { height: "12px", opacity: "0.5" },
          "75%": { height: "20px", opacity: "1" },
        },
      },
      animation: {
        // 动画持续时间和无限循环
        loading: "loading 3s ease-in-out infinite",
      },
    },
  },
  plugins: [require("@tailwindcss/forms")],
};
export default config;
