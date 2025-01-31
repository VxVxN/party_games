import {useQuery} from "@tanstack/react-query";
import {axiosInstance} from "./axiosInstance.ts";

async function getTopics() {
    const data = await axiosInstance.get<string[]>("/topic/list");

    return data.data;
}

export const useTopicList = () => useQuery({queryKey: ['topics'], queryFn: getTopics})