import dayjs from "dayjs";
const url = process.env.NEXT_PUBLIC_URL;
// const url = "https://f1f5-183-13-202-191.ngrok-free.app/api";

export async function GET() {
  const response = await fetch(`${url}/api/v1/nfts/latest/10`, {
    method: "GET",
    cache: "no-store",
  });
  const data = await response.json();
  if (data.count > 0) {
    data.data.forEach((item: any) => {
      item.CreatedAt = dayjs(item.CreatedAt).format("YYYY-MM-DD");
      item.UpdatedAt = dayjs(item.UpdatedAt).format("YYYY-MM-DD");
    });
  }
  return Response.json(data);
}
