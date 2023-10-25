async function fetchLocalJSON() {
  try {
    //   const response = await fetch('js/data.json',{mode: 'no-cors'});

    const response = await fetch('json/data.json', {
      method: 'GET',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
    });
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }

    return response;
  } catch (error) {
    console.error('There has been a problem with your fetch operation:', error);
  }
}

async function getTodayTime() {
  const currentDate = new Date();
  const year = currentDate.getFullYear();
  const month = (currentDate.getMonth() + 1).toString().padStart(2, '0');
  const day = currentDate.getDate().toString().padStart(2, '0');
  let todayTime = `${year}-${month}-${day}`;

  return todayTime;
}

async function getNeddedTimeJSON(response, formattedDate) {
  try {
    data = await response.json();

    return data.result.find(entry => entry.Date === formattedDate);
  } catch (error) {
    console.error('There has been a problem with your fetch operation:', error);
  }
}

async function setPrayerTime(todayTime) {
  try {
    const prayers = ['fajr', 'sunrise', 'dhuhr', 'asr', 'maghrib', 'isha'];

    if (todayTime) {
      for (i = 0; i < 6; i++) {
        document.getElementsByClassName("time_namaz")[i].getElementsByTagName("span")[0].innerHTML = todayTime[prayers[i]];
      }
    } else {
      throw new Error('Can\'t read JSON File.');
    }
  } catch (error) {
    console.error('Error fetching JSON:', error);
  }
}

// Determine which prayer now
async function whichPrayerNext(todayTimeJSON) {
  const prayers = ['fajr', 'sunrise', 'dhuhr', 'asr', 'maghrib', 'isha'];
  let currentDate = new Date();
  let hoursNow = currentDate.getHours();
  let minutesNow = currentDate.getMinutes();

  for (i = 0; i < 6; i++) {
    let dateTime = todayTimeJSON[prayers[i]].split(':');
    const h = parseInt(dateTime[0], 10);
    const m = parseInt(dateTime[1], 10);

    if (hoursNow == h) {
      if (minutesNow <= m) {
        return dateTime[0] + ':' + dateTime[1];
      }
    } else if (hoursNow < h) {
      return dateTime[0] + ':' + dateTime[1];
    }
  }

  dateTime = todayTimeJSON['fajr'].split(':');
  const h = parseInt(dateTime[0], 10);
  const m = parseInt(dateTime[1], 10);
  return dateTime[0] + ':' + dateTime[1];
}

async function setActiveMode(time) {
  for (i = 0; i <= 6; i++) {
    if (document.getElementsByClassName("time_namaz")[i].getElementsByTagName("span")[0].innerHTML == time) {
      if (i == 1) {
        document.getElementsByClassName("time_namaz")[i - 1].classList.add('active');
        document.getElementsByClassName("time_namaz")[5].classList.remove('active');
        return;
      }
      if (i == 0) {
        document.getElementsByClassName("time_namaz")[5].classList.add('active');
        document.getElementsByClassName("time_namaz")[4].classList.remove('active');
        return;
      }
      document.getElementsByClassName("time_namaz")[i - 1].classList.add('active');
      document.getElementsByClassName("time_namaz")[i - 2].classList.remove('active');
      return;
    }
  }
}

async function updateTimer(formattedDate, formattedTime, stopInterval) {
  // Set the target date and time for the countdown
  let date = formattedDate.split('-'); // Split the date portion
  let time = formattedTime.split(':'); // Split the time portion
  const year = parseInt(date[0], 10);
  const month = parseInt(date[1], 10) - 1; // Adjust month to be 0-based
  let day = parseInt(date[2], 10);
  const hours = parseInt(time[0], 10);
  const minutes = parseInt(time[1], 10);

  if (formattedTime == document.getElementById("tan").innerHTML) {
    day += 1;
  }
  const targetDate = new Date(year, month, day, hours, minutes, 0).getTime();
  const currentDate = new Date().getTime();
  const timeRemaining = targetDate - currentDate;
  
  if (timeRemaining <= 1) {
    clearInterval(stopInterval);
    go();
  }

  const hoursOutput = Math.floor((timeRemaining % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
  const minutesOutput = Math.floor((timeRemaining % (1000 * 60 * 60)) / (1000 * 60));
  const secondsOutput = Math.floor((timeRemaining % (1000 * 60)) / 1000);

  const timerText = `${hoursOutput.toString().padStart(2, '0')}:${minutesOutput.toString().padStart(2, '0')}:${secondsOutput.toString().padStart(2, '0')}`;
  document.getElementById('time_remain').getElementsByTagName("span")[0].innerHTML = timerText;
  if (hoursOutput == 0 && (minutesOutput < 20)) {
    document.getElementById('time_remain').parentElement.classList.add('bg-danger');
  } else if (hoursOutput == 0 && (minutesOutput < 40)) {
    document.getElementById('time_remain').parentElement.classList.add('bg-warning');
  } else {
    document.getElementById('time_remain').parentElement.classList.add('bg-success');
  }
}
let formattedDate = "hello";
let prayerNext;
let timerInterval;

async function go() {
  const response = await fetchLocalJSON();
  formattedDate = await getTodayTime();
  let todayTimeJSON = await getNeddedTimeJSON(response, formattedDate);
  await setPrayerTime(todayTimeJSON);
  prayerNext = await whichPrayerNext(todayTimeJSON);
  await setActiveMode(prayerNext);
  // Initial check of how much time remains
  // await updateTimer();

  // Set the timer to update every second
  stopInterval = setInterval(() => {
    updateTimer(formattedDate, prayerNext, stopInterval);
  }, 1000);
}
go();