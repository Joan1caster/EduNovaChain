import { cookies } from "next/headers";

const url = process.env.NEXT_PUBLIC_URL;
// const url = "https://f1f5-183-13-202-191.ngrok-free.app/api";

export async function GET(wallet: `0x${string}`) {
  const cookie = cookies();
  const token = cookie.get("token");
  if (token) {
    return Response.json("");
  } else {
    const response = await fetch(`${url}/api/v1/siweMessage?wallet=${wallet}`, {
      method: "GET",
      cache: "no-store",
      headers: {
        "ngrok-skip-browser-warning": "off",
      },
    });
    const data = await response.json();
    return Response.json(data);
  }
}
