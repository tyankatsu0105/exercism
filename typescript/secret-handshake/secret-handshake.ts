/**
 * 10000: 'reverse the order of the operations in the secret handshake'
 */
const secrets = {
  1:  'wink',
  10:  'double blink',
  100:  'close your eyes',
  1000:  'jump'
} as const

class HandShake {
  private binary = ''

  constructor(private input: number){
    this.binary = this.toBinary(input)
  }

  public commands(): string[] {
    let result = []
    const splittedBinary = [...this.binary].reverse()
    
    if(splittedBinary[0] && splittedBinary[0] === '1') result.push(secrets[1])
    if(splittedBinary[1] && splittedBinary[1] === '1') result.push(secrets[10])
    if(splittedBinary[2] && splittedBinary[2] === '1') result.push(secrets[100])
    if(splittedBinary[3] && splittedBinary[3] === '1') result.push(secrets[1000])
    if(splittedBinary[4] && splittedBinary[4] === '1') result = result.reverse()
    
    return result
  }

  private toBinary(input: number): string {
    return input.toString(2)
  }
}

export default HandShake