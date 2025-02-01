import { Link, useParams } from "react-router";
import Questions from "../components/Questions";
import PageHeader from "../components/PageHeader";

function GamePage() {
  const { topic } = useParams<{ topic: string }>();

  return (
    <div className="page">
      <PageHeader title={topic || ""} />
      {topic ? <Questions topics={[topic]} /> : <div>Empty</div>}
      <Link
        to={"/"}
        className="absolute bottom-5 right-5 p-3 bg-stone-900 bg-opacity-80 text-rose-200 rounded-xl font-medium"
      >
        Go to menu
      </Link>
    </div>
  );
}

export default GamePage;
