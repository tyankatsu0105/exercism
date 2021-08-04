type ExpectedStr = `${number}` | ' ' | '\n'
type AssertExpectedStr<T extends string> = T extends '' ? true : T extends `${ExpectedStr}${infer Rest}` ? AssertExpectedStr<Rest> : false

/**
 * Expected numbers as string
 * @example
 * const str = '1 2 3 \n4 5 6\n7'
 * const result: NumbersAsString<typeof str> = str
 * ---
 * const str = '1 2 3 \n4 a' // error
 */
type NumbersAsString<T extends string> = AssertExpectedStr<T> extends true ? T : never

class Matrix<T extends string> {
  public rows: number[][]
  public columns: number[][]

  constructor(private enumNumbers: NumbersAsString<T>) {
    this.rows = this.getRows()
    
    this.columns = this.getColumns()
  }

  private getRows(): number[][] {
    const rows = this.enumNumbers.split('\n')

    return rows.map(row => row.split(' ').map(Number))
  }

  private getColumns(): number[][] {
    return this.rows.reduce<number[][]>((acc, current) => {
      current.forEach((value, index) => {
        if(!acc[index]) acc[index] = []

        acc[index].push(value)
      })

      return acc
    }, [])
    
  }
}

export default Matrix