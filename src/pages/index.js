import React from'react';
import Link from'next/link';
	
export default class extends React.Component {
	render() {
		return ( 
		
      <div>
        <title>Backend Task</title>
        <h1>Backend Task</h1>
        <div className='form'>
          <div><input placeholder='enter id'></input>
            <br></br>
            <input placeholder='enter text'></input>
            <br></br>
            <button>Post!</button>
          </div>
          <div><button>Get all!</button></div>
          <div>
            <input placeholder='enter id'></input>
            <br></br>
            <button>Delete!</button>
          </div>
          </div>
		</div>
		)
	}
}
