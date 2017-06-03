<div class="container">
  <div class="row">
    <div class="hero-text">
     <h1>Current Prices</h1>
   </div>
 </div>
</div>
</div>
{{.starttime}}
<div class="container">
  <div class="row">
    <table class="table">
      <thead>
        <tr>
          <th>Id</th>
          <th>Start Time</th>
          <th>End Time</th>
          <th>Average Price (cents/kWh)</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        {{range $record := .records}}
        <tr>
          <td>{{$record.Id}}</td>
          <td>{{$record.Starttime}}</td>
          <td>{{$record.Endtime}}</td>
          <td>{{$record.Avgprice}}</td>
        </tr>
        {{end}}
      </tbody>
    </table>
  </div>
</div>