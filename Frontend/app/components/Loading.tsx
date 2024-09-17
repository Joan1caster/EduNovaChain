type LoadingProps = {
  show?: boolean;
  fullScreen?: boolean;
};
export default function Loading({
  show = false,
  fullScreen = true,
}: LoadingProps) {
  return (
    <div
      className={`${show ? "visible" : "invisible"} ${fullScreen ? "fixed w-screen h-screen" : "absoulte w-full h-full"} top-0 left-0 z-50 flex items-center justify-center`}
    >
      <div className="flex w-14 h-24 relative space-x-2">
        <div className="w-2.5 h-2.5 absolute left-2 bottom-0 bg-blue-500 rounded-full animate-[loading_1s_ease-in-out_infinite]"></div>

        <div className="w-2.5 h-2.5 absolute left-4 bottom-0 bg-blue-500 rounded-full animate-[loading_1s_ease-in-out_infinite] [animation-delay:0.2s]"></div>

        <div className="w-2.5 h-2.5 absolute left-8 bottom-0 bg-blue-500 rounded-full animate-[loading_1s_ease-in-out_infinite] [animation-delay:0.4s]"></div>

        <div className="w-2.5 h-2.5 absolute left-12 bottom-0 bg-blue-500 rounded-full animate-[loading_1s_ease-in-out_infinite] [animation-delay:0.6s]"></div>
      </div>
    </div>
  );
}
