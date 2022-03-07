
// Get the HTML canvas by its id
plots = document.getElementById("plots");


var datas = [90, 230, 80, 81, 56, 55, 80];

// Example datasets for X and Y-axes
var months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul"]; //Stays on the X-axis
var traffic = [65, 59, 80, 81, 56, 55, 60] //Stays on the Y-axis
var empt = []
for (var i=0; i<datas.length;  i++) {
    empt.push(Math.round((datas[i]/datas.reduce((a, b) => a + b, 0))*360))
    
}

// Create an instance of Chart object:
new Chart(plots, {
  type: 'pie', //Declare the chart type
  data: {
    labels: months, //Defines each segment
    datasets: [{
      data: traffic, //Determines segment size
      //Color of each segment
      backgroundColor: [
                  "rgba(196, 190, 183)",
                  "rgba(21, 227, 235)",
                  "rgba(7, 150, 245)",
                  "rgba(240, 5, 252)",
                  "rgba(252, 5, 79)",
                  "rgb(0,12,255)",
                  "rgb(17, 252, 5)"],
    }]
  },
  options:{
    legend: {display: true}, //This is true by default.
  
  }
  
});

// options:{
//   legend:{display: false},
// }


// new Chart(plots, {
//     type: 'line',
  
//   data: {
//     labels: mon,
//     datasets: [{
//     data: traffic,
      
//     //   pointBackgroundColor: "rgb(0,12,255)",
//       backgroundColor: [
//           "rgba(196, 190, 183)",
//           "rgba(21, 227, 235)",
//           "rgba(7, 150, 245)",
//           "rgba(240, 5, 252)",
//           "rgba(252, 5, 79)",
//           "rgb(0,12,255)",
//           "rgb(17, 252, 5)"],
//           borderColor: "rgb(17, 252, 5)",
      
//       fill: false,
//       gridLines: false


//     },
// ]
//   },
//   options: {
//     legend: {display: false},
//     scales: {
        
//       yAxes: [{ticks: {min: 50}}],
//     }
//   }
// });