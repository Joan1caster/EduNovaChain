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
    },
  },
  plugins: [require("@tailwindcss/forms")],
};
export default config;
