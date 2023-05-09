import { useState } from 'react'
import './App.css'

interface Record {
  ID: string
  Title: string
  Artist: string
  Price: number
}

type Data = Record[]

interface TableProps {
  data: Data;
}

const TableComponent: React.FC<TableProps> = ({ data }) => {
  console.log(data);
  return (
    <table>
      <thead>
        <tr>
          <th>ID</th>
          <th>Title</th>
          <th>Artist</th>
          <th>Price</th>
        </tr>
      </thead>
      <tbody>
        {data.map((item) => (
          <tr key={item.ID}>
            <td>{item.ID}</td>
            <td>{item.Title}</td>
            <td>{item.Artist}</td>
            <td>{item.Price}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}

function ButtonAndTable() {
  const [data, setData] = useState([]);
  const fetchData = async () => {
    const response = await fetch("http://localhost:8081/albums");
    const jsonData = await response.json();
    setData(jsonData);
    console.log(jsonData);
  };
  return (
    <div>
      <button onClick={fetchData}>Fetch Data</button>
      <TableComponent data={data} />
    </div>
  );
}
function App() {
  return (
    <><ButtonAndTable /></>
  )
}
export default App
