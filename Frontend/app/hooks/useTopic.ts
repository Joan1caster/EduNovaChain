"use client";

import { useAsyncEffect } from "ahooks";
import { useState } from "react";

export default function useTopic() {
  const [currentGrade, setCurrentGrade] = useState<undefined | number>();
  const [gradeList, setGradeList] = useState([]);

  const [currentSubject, setCurrentSubject] = useState();
  const [subjectList, setSubjectList] = useState([]);

  const [topicList, setTopicList] = useState([]);

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
    }
  }, [currentGrade]);

  return {
    currentGrade,
    gradeList,

    currentSubject,
    subjectList,
    topicList,
  };
}
