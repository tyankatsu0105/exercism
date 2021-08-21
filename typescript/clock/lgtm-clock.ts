// https://exercism.io/tracks/typescript/exercises/clock/solutions/b17dd29303db41a4a7542a6d0b7b5e95

class Clock {
  readonly minutes: number

  constructor(hour: number, minute: number = 0) {
      this.minutes = hour * 60 + minute
  }

  private get hour(): number {
      const h = Math.floor(this.minutes / 60 % 24)
      return h >= 0 ? h : 24 + h
  }

  private get minute(): number {
      const m = this.minutes % 60
      return m >= 0 ? m : 60 + m
  }

  private padZero(n: number): string {
      return n.toString().padStart(2, '0')
  }

  toString(): string {
      return `${this.padZero(this.hour)}:${this.padZero(this.minute)}`
  }

  plus(minute: number): Clock {
      return new Clock(0, this.minutes + minute)
  }

  minus(minute: number): Clock {
      return new Clock(0, this.minutes - minute)
  }

  equals(other: Clock): boolean {
      return this.toString() === other.toString()
  }
}

export default Clock