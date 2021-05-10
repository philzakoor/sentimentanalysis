import React from "react";
import ReactWordcloud from 'react-wordcloud';

const Results = (props) => {
    if (props.results == null) return (
             <div className='Results'>
            <p>Type to search</p>
        </div>
    );
    const {results}  = props.results;
    if (!results || results.length === 0) {return <p>No Tweets Found</p>;}
    return (
        <div className="Results">
        <h2>Sentiment Results</h2>
        {results.map((results) => {
            return (
                <span key={results.Filter}> {results.Sentiment}
                <ReactWordcloud words = {results.FrequencyWords} /></span>
            );
        })}
        </div>
    );

}

export default Results;