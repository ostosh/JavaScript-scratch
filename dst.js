function isDaylightSavings(timestamp){
  var date = new Date(timestamp);
  var dateBegYear = new Date(date.getFullYear(), 0, 1);
  var dateMidYear = new Date(date.getFullYear(), 6, 1);
  return date.getTimeZoneOffset() < Math.max(dateBegYear.getTimeZoneOffset(), dateMidYear.getTimeZoneOffset());
}
