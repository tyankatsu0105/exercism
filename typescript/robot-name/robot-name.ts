export default class Robot {
  savedName = ''
  usedNames: Set<string> = new Set()

  constructor() {
    this.savedName = this.generateName()
  }

  public get name(): string {
    return this.savedName
  }

  private generateName(): string {
    const generatedName = `${this.generateRandomCharacters('ABCDEFGHIJKLMNOPQRSTUVWXYZ', 2)}${this.generateRandomCharacters('0123456789', 3)}`

    if (this.usedNames.has(generatedName)) {
      return this.generateName()
    }

    this.usedNames.add(generatedName)  

    return generatedName
  }

  private generateRandomCharacters(character: string, length: number = 1): string {
    let result = ''

    for (let index = 0; index < length; index++) {
      result += character[Math.floor(Math.random() * character.length)]
    }

    return result
  }

  public resetName(): void {
    this.savedName = this.generateName()
  }

  public static releaseNames(): void {
    console.log('release');
    
  }
}
