import React, { Component } from 'react';

class Filter extends Component {
    constructor(props) {
        super(props);
        this.state ={
            filter: ''
        };
    }

    _start = (e) => {
        if (e.key === 'Enter') {
            alert('enter')
        }
    }

    render() {
        return (
          <form>
              <input type="text" name="filter" onKeyDown={this._start} />
              <button  onClick={this._start} />
          </form>
        );

    }
}