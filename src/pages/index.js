import React from'react';
import Link from'next/link';
	
export default class extends React.Component {
	render() {
		return ( 
		
      <div>
        <title>Backend Task</title>
        <h1>Backend Task</h1>
        <div className='form'>
          <div>
            <h2>Get All</h2>
            <button>Get All!</button></div>
          <div>
          <h2>Post</h2>
            <input placeholder='enter id'></input>
            <br></br>
            <input placeholder='enter text'></input>
            <br></br>
            <button>Post!</button>
          </div>
          <div>
          <h2>Update</h2>
            <input placeholder='enter id'></input>
            <br></br>
            <input placeholder='enter text'></input>
            <br></br>
            <button>Update!</button>
          </div>
          <div>
          <h2>Delete</h2>
          <br></br>
            <input placeholder='enter id'></input>
            <br></br>
            <button>Delete!</button>
            </div>
          </div>
		</div>
		)
	}
}
