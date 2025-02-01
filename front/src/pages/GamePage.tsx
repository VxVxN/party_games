import {useParams} from "react-router";
import Questions from "../components/Questions";
import PageHeader from "../components/PageHeader";

function GamePage() {
    const {topic} = useParams<{ topic: string }>();

    return <div className='page'>
        <PageHeader title={topic || ''}/>
        {
            topic ? <Questions topics={[topic]}/> : <div>Empty</div>
        }
    </div>
}

export default GamePage;