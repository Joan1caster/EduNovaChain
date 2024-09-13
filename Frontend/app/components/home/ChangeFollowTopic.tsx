"use client";

import { TagType } from "@/app/types";
import { useState, useEffect, useReducer } from "react";
import { useAsyncEffect } from "ahooks";

type Props = {
  followTopics: TagType[];
  onUpdateFollowTopics: (topics: TagType[]) => void;
};

type InitialState = {
  topicKeys: number[];
  topics: TagType[];
};
type Action =
  | { type: "init"; payload: TagType[] }
  | { type: "add"; payload: TagType }
  | { type: "remove"; payload: TagType };

const initialState: InitialState = {
  topicKeys: [],
  topics: [],
};

function reducer(state: InitialState, action: Action) {
  switch (action.type) {
    case "init": {
      state.topicKeys = action.payload.map((item) => item.id);
      state.topics = action.payload.slice();
      return state;
    }
    case "add": {
      return {
        topicKeys: [...state.topicKeys, action.payload.id],
        topics: [...state.topics, action.payload],
      };
    }
    case "remove": {
      const position = state.topicKeys.indexOf(action.payload.id);
      state.topicKeys.splice(position, 1);
      state.topics.splice(position, 1);
      return Object.assign({}, state);
    }
  }
}

export default function ChangeFollowTopic({
  followTopics,
  onUpdateFollowTopics,
}: Props) {
  const [state, dispatch] = useReducer(reducer, initialState);
  const [currentGrade, setCurrentGrade] = useState<number | null>();
  const [gradeList, setGradeList] = useState<TagType[]>([]);

  const [currentSubject, setCurrentSubject] = useState<number | null>();
  const [subjectList, setSubjectList] = useState<TagType[]>([]);

  const [topicList, setTopicList] = useState<TagType[]>([]);

  useAsyncEffect(async () => {
    const response = await (await fetch("/api/grade")).json();
    if (response.count > 0) {
      setGradeList(response.data);
      setCurrentGrade(response.data[0].id);
    }
  }, []);

  useAsyncEffect(async () => {
    if (currentGrade) {
      const response = await (
        await fetch(`/api/subject?id=${currentGrade}`)
      ).json();
      if (response.count > 0) {
        setSubjectList(response.data)
        setCurrentSubject(response.data[0].id)
      }
    }
  }, [currentGrade]);

  const onChangeGrade = (grade: TagType) => {
    setCurrentGrade(grade.id);
  };
  const onChangeSubject = (subject: TagType) => {
    setCurrentSubject(subject.id);
  };

  useEffect(() => {
    dispatch({ type: "init", payload: followTopics });
  });
  return (
    <div>
      <div className="flex justify-between items-center">
        <div className="text-xl">我关注的主题</div>
        <button
          className="px-4 py-1 bg-primary rounded-sm text-sm text-white"
          onClick={() => onUpdateFollowTopics(state.topics)}
        >
          完成
        </button>
      </div>
      <div className="flex my-4">
        {/* choiced topics start */}
        <div className="w-1/3 *:inline-block gap-2 *:text-sm *:font-light *:py-0.5 *:px-3 *:mr-2 *:mb-4 *:border *:rounded-full">
          {state.topics.map((item) => (
            <div
              key={item.id}
              className={
                item.id === 0
                  ? "border-primary/50 text-primary/50"
                  : "flex items-center gap-2 border-primary text-primary"
              }
              onClick={() =>
                item.id !== 0 && dispatch({ type: "remove", payload: item })
              }
            >
              <span>{item.name}</span>
              {item.id !== 0 && (
                <svg
                  t="1725173545826"
                  className="inline-block mb-0.5 ml-1"
                  viewBox="0 0 1024 1024"
                  version="1.1"
                  xmlns="http://www.w3.org/2000/svg"
                  p-id="1692"
                  width="10"
                  height="10"
                >
                  <path
                    d="M557.12 512l313.6-313.28a32.128 32.128 0 1 0-45.44-45.44L512 466.88l-313.28-313.6a32.128 32.128 0 0 0-45.44 45.44L466.88 512l-313.6 313.28a32 32 0 0 0 0 45.44 32 32 0 0 0 45.44 0L512 557.12l313.28 313.6a32 32 0 0 0 45.44 0 32 32 0 0 0 0-45.44L557.12 512z"
                    fill="#1474FC"
                    p-id="1693"
                  ></path>
                </svg>
              )}
            </div>
          ))}
        </div>
        {/* choiced topics end */}

        {/* all topics start */}
        <div className="pl-4 border-l border-l-primary-light_bg/50">
          {/* grade start */}
          <div>
            <span className="text-base pr-4">年级</span>
            {gradeList.map((item) => (
              <span
                key={item.id}
                className={`${currentGrade === item.id ? "text-primary bg-primary-light_bg/50" : ""} inline-block py-1 px-4 rounded-full text-sm font-light text-gray-500 cursor-pointer hover:bg-primary-light_bg/50`}
                onClick={() => onChangeGrade(item)}
              >
                {item.name}
              </span>
            ))}
          </div>
          {/* grade end */}
          {/* subject start */}

          <div className="my-4">
            <span className="text-base pr-4">学科</span>
            {subjectList.map((item) => (
              <span
                key={item.id}
                className={`${currentSubject === item.id ? "text-primary bg-primary-light_bg/50" : ""} inline-block py-1 px-4 rounded-full text-sm  font-light text-gray-500 cursor-pointer hover:bg-primary-light_bg/50`}
                onClick={() => onChangeSubject(item)}
              >
                {item.name}
              </span>
            ))}
          </div>
          {/* sugject end */}
          {/* topic start */}
          <div className="flex flex-wrap gap-4">
            {topicList.map((item) => (
              <>
                {state.topicKeys.includes(item.id) ? (
                  <div
                    key={item.id}
                    className="py-1 px-4 rounded-full text-sm font-light bg-primary text-white"
                  >
                    {item.name}
                  </div>
                ) : (
                  <div
                    key={item.id}
                    className="py-1 px-2 rounded-full text-sm font-light border border-primary text-primary cursor-pointer"
                    onClick={() => dispatch({ type: "add", payload: item })}
                  >
                    <svg
                      t="1725175223714"
                      className="inline-block mb-0.5 mr-1"
                      viewBox="0 0 1024 1024"
                      version="1.1"
                      xmlns="http://www.w3.org/2000/svg"
                      p-id="1841"
                      width="12"
                      height="12"
                    >
                      <path
                        d="M892.16 480H544V131.84a32 32 0 0 0-64 0V480H131.84a32 32 0 0 0 0 64H480v348.16a32 32 0 1 0 64 0V544h348.16a32 32 0 1 0 0-64z"
                        fill="#1474FC"
                        p-id="1842"
                      ></path>
                    </svg>
                    {item.name}
                  </div>
                )}
              </>
            ))}
          </div>
          {/* topic end */}
        </div>
        {/* all topics end */}
      </div>
    </div>
  );
}
