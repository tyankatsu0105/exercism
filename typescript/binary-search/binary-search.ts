type Array = number[]

export default class BinarySearch {
  public array: Array | undefined = undefined

  constructor(private input: Array) {
    if(this.isSorted(input)) this.array = input
  }

  private isSorted(array: Array): boolean {
    for (let index = 0; index < array.length - 1; index++) {
      if(array[index] > array[index + 1]) return false
    }

    return true
  }

  public indexOf(num: number): number {
    if(!this.array) return -1

    let left = 0
    /**
     * [0,1,2,3,4,5] length
     * expected ... 5
     * recieved 6
     */
    let right = this.array.length - 1
    let half = 0

    while (left < right) {
      half = Math.ceil((right + left) / 2)

      if(num === this.array[half]) return half

      if(num < this.array[half]) {
        right = half
      } else if(num > this.array[half]) {
        left = half
      } else {
        break
      }
    }

    return -1
  }
}
