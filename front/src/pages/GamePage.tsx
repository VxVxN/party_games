import {useParams} from "react-router";
import Questions from "../components/Questions";
import PageHeader from "../components/PageHeader";

function GamePage() {
    const params = useParams<{ categoryId: string }>();

    return <div className='page'>
        <PageHeader title={`game ${params.categoryId}`}/>
        <Questions/>
    </div>
}

export default GamePage;