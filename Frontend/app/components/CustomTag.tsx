export const AiTag = () => {
  return (
    <div className="w-6 h-4 inline-block bg-[url('/images/slice/ai_bg.jpg')] bg-contain bg-no-repeat text-center text-sm text-white">
      AI
    </div>
  );
};

export const OrderTag = ({
  order,
  bg = true,
}: {
  order: number;
  bg?: boolean;
}) => {
  return (
    <div
      className={`w-6 h-6 inline-block bg-[url('/images/slice/${order > 3 && bg ? "n.jpg" : order + ".png"}')] bg-contain bg-no-repeat bg-opacity-80 text-center text-sm ${bg ? "text-[#333]" : "text-[#999]"} leading-6`}
    >
      {order}
    </div>
  );
};
