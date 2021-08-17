type Student = string
type Grade = string
type StudentMap = Map<Grade | number, Student[]>
type GradeMap = Map<Student, Grade>

type ExtractMap<M extends Map<any, any>> = M extends Map<infer K, infer V> ? {
  key: K,
  value: V
} : never

class GradeSchool {
  public students: StudentMap = new Map()
  private grades: GradeMap = new Map()

  constructor() {}

  public studentRoster(): StudentMap {
    return this.toDeepCloneMap(this.students)
  }

  public studentsInGrade(grade: Grade | number): ExtractMap<StudentMap>['value'] {
    return this.toDeepCloneMap(this.students).get(this.toGrade(grade)) || []
  }

  public addStudent(student: Student, grade: Grade | number): void {
    if (this.isExistStudentAtDifferentGrade(student, this.toGrade(grade))) {
      this.deleteDuplicateStudent(student, this.toGrade(grade))
      
      return
    }
    
    const existValue = this.students.get(this.toGrade(grade)) || []
    
    this.students.set(this.toGrade(grade), [...existValue, student].sort())
    this.grades.set(student, this.toGrade(grade))
  }

  private toDeepCloneMap<T extends Map<unknown, unknown>>(map: T): T {
    return new Map(JSON.parse(JSON.stringify(Array.from(map)))) as T
  }

  private deleteDuplicateStudent(student: Student, grade: Grade): void {
    const existStudentGrade = this.grades.get(student)
    if(!existStudentGrade) return
    const existValue = this.students.get(existStudentGrade)
    if(!existValue) return
    
    const index = existValue.indexOf(student)
    
    this.students.set(grade, existValue.splice(index, 1))
    this.grades.delete(student)
  }

  private isExistStudentAtDifferentGrade(student: ExtractMap<GradeMap>['key'], grade: ExtractMap<GradeMap>['value']): boolean {
    if(this.grades.get(student) === undefined) return false

    return this.grades.get(student) !== this.toGrade(grade)
  }

  private toGrade(grade: Grade | number): Grade {
    return String(grade)
  }
}

export default GradeSchool