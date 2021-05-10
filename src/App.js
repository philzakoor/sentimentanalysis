import './App.css';
import {Component} from "react";
import Results from './components/Results';
import WithResultsLoading from "./components/WithResultsLoading";
import axios from "axios";

class App extends Component {
    constructor(props) {
        super(props);
        this.state ={
            filter: '',
            results: null,
            loading: false
        };
    }

    _startByKey = (e) => {
        if (e.key === 'Enter') {
            this.fetchAnalysis().then(r =>
                this.setState({loading: false, results: r}));
        }
    }

    _startByClick = () => {
        this.fetchAnalysis().then(r =>
            this.setState({loading: false, results: r.data}));
    }

    _setFilter = (e) => {
        this.setState({filter: e.target.value, loading: true})
    }

    async fetchAnalysis() {
        const apiUrl = 'http://localhost:8080/sentiment/' + this.state.filter;
        await axios.get(apiUrl);
    }

    render() {
        const ResultLoading = WithResultsLoading(Results);
        return (
            <div className="App">
                <header className="App-header">
                    <h1>Twitter Sentiment Analysis.</h1>
                    <form>
                        <input type="text" id="filter" onChange={this._setFilter} onKeyDown={this._startByKey}/>
                        <button onClick={this._startByClick}>Analyze</button>
                    </form>
                </header>
                <div className="Results">
                    <ResultLoading isLoading={this.state.loading} results={this.state.results} />
                </div>
            </div>
        );

    }
}

export default App;

