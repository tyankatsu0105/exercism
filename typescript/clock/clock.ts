type Hour = '00' | '01' | '02' | '03' | '04' | '05' | '06' | '07' | '08' | '09' | '10' | '11' | '12' | '13' | '14' | '15' | '16' | '17' | '18' | '19' | '20' | '21' | '22' | '23'

type Minute = '00' | '01' | '02' | '03' | '04' | '05' | '06' | '07' | '08' | '09' | '10' | '11' | '12' | '13' | '14' | '15' | '16' | '17' | '18' | '19' | '20' | '21' | '22' | '23' | '24' | '25' | '26' | '27' | '28' | '29' | '30' | '31' | '32' | '33' | '34' | '35' | '36' | '37' | '38' | '39' | '40' | '41' | '42' | '43' | '44' | '45' | '46' | '47' | '48' | '49' | '50' | '51' | '52' | '53' | '54' | '55' | '56' | '57' | '58' | '59'

type Time = `${Hour}:${Minute}`

class Clock {
  private hours = 0
  private minutes = 0

  constructor(private initialHour: number, private initialMinute: number = 0) {
    this.hours = this.initialHour
    this.minutes = this.initialMinute
  }

  private get hour(): number {
    let result: number = this.hours
    result += Math.floor(this.minutes / 60)

    if(result > 0) result = result % 24
    if(result < 0) result = result % 24 + 24
    if(result === 24) result = 0

    return result
  }

  private get minute(): number {
    let result: number = this.minutes
    
    if(result > 0) result = result % 60
    if(result < 0) result = result % 60 + 60
    if(result === 60) result = 0

    return result
  }

  public toString(): Time {
    return `${this.toFormattedTimeValue(this.hour)}:${this.toFormattedTimeValue(this.minute)}` as Time
  }

  public equals(clock: Clock): boolean {
    return this.toString() === clock.toString()
  }

  public plus(minutes: number): Clock {
    this.minutes += minutes
    return this
  }

  public minus(minutes: number): Clock {
    this.minutes -= minutes
    return this
  }

  private toFormattedTimeValue(num: number): string {
    return num.toString().padStart(2, '0')
  }
}

export default Clock

// class Clock {
//   hour = 0
//   minute = 0

//   constructor(private initialHour: number, private initialMinute?: number) {
//     this.hour = this.initialHour
//     this.minute = this.initialMinute || 0
//   }

//   public toString(): Time {
//     let minute = this.minute

//     if(minute > 0) minute = minute % 60
//     if(minute < 0) minute = minute % 60 + 60
//     if(minute === 60) minute = 0

//     let hour = this.hour
//     hour += Math.floor(this.minute / 60)
    
//     if(hour > 0) hour = hour % 24
//     if(hour < 0) hour = hour % 24 + 24
//     if(hour === 24) hour = 0

//     return `${this.toFormattedTimeValue(hour)}:${this.toFormattedTimeValue(minute)}` as Time
//   }

//   public equals(clock: Clock): boolean {
//     return this.toString() === clock.toString()
//   }

//   public plus(minutes: number): Clock {
//     this.minute += minutes
//     return this
//   }

//   public minus(minutes: number): Clock {
//     this.minute -= minutes
//     return this
//   }

//   private toFormattedTimeValue(num: number): string {
//     return num.toString().padStart(2, '0')
//   }
// }

// export default Clock