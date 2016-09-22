DECLARE 
  @start datetime = '1970-01-01T00:00:00.000',
  @end datetime = SWITCHOFFSET(SYSDATETIMEOFFSET(), '-00:00'),
  @dayInMills bigint = 60 * 60 * 24 * 1000,
  @unixTimestamp varchar(13)
;

SET @
  unixTimestamp = (@dayInMills * DATEDIFF(DAY, @start, @end)) - 
    DATEDIFF(millisecond, @end, CAST(@end AS date));

-- select unix timestamp in mills
SET @unixTimestamp = (SELECT
  RIGHT(CONCAT('000', @unixTimestamp), LEN(@unixTimestamp) + 3));
  
DECLARE @start datetime = '1970-01-01T00:00:00.000',
  	@dayInMills bigint = 60 * 60 * 24 * 1000,
    @unixTimestamp bigint
;

-- select datetime
SELECT
  DATEADD(millisecond, @unixTimestamp % @dayInMills,
  DATEADD(DAY, @unixTimestamp / @dayInMills, @start)
  )
