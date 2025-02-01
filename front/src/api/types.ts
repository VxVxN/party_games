export type RequestTopicRecords = {
  topics: string[];
  page: number;
  page_size?: number;
};

export type ResponseTopicRecords = {
  count_page: number;
  records: string[];
};
