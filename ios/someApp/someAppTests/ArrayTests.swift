
import XCTest
@testable import someApp

class ArrayTests: XCTestCase {
    
  func testToDictionary() {
    let array = ["1", "2", "3"]
    XCTAssertEqual(array.toDictionary { item in item }, ["1": "1", "2": "2", "3": "3"])
    XCTAssertEqual(array.toDictionary(keySelector: { item in item }) { item in String(format: "%@%@", item, item) }, ["1": "11", "2": "22", "3": "33"])
  }

  func testToMultiDictionary() {
    let array = ["1", "22", "3", "444", "55"]
    var dict = array.toMultiDictionary { item in item.characters.count}
    XCTAssertEqual(dict.count, 3)
    XCTAssertEqual(dict[1]!, ["1", "3"])
    XCTAssertEqual(dict[2]!, ["22", "55"])
    XCTAssertEqual(dict[3]!, ["444"])
    dict = array.toMultiDictionary(keySelector: { item in item.characters.count }) { item in
      String(format: "%@%@", item, item)
    }
    XCTAssertEqual(dict.count, 3)
    XCTAssertEqual(dict[1]!, ["11", "33"])
    XCTAssertEqual(dict[2]!, ["2222", "5555"])
    XCTAssertEqual(dict[3]!, ["444444"])
  }
    
}
