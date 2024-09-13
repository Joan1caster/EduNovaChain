import { NextRequest } from "next/server";

const url = process.env.NEXT_PUBLIC_URL;
// const url = "https://f1f5-183-13-202-191.ngrok-free.app/api";

export async function GET(request: NextRequest) {
  const id = request.nextUrl.searchParams.get("id");
  const response = await fetch(`${url}/api/v1/subject/${id}`, {
    method: "GET",
    cache: "no-store",
  });
  const data = await response.json();
  data.data.forEach((item) => {
    item.id = item.ID;
    item.name = item.Name;
  });
  return Response.json(data);
}
