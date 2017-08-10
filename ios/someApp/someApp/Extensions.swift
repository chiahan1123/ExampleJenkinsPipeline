
import Foundation

// MARK: - UI Extensions

// MARK: - Misc

extension Array {

  /// Transforms an array to a dictionary with each key-value pair being a key from the key selector and the value as the element itself respectively. Inspired by RxJava's toMap.
  ///
  /// - Parameter keySelector: Target key selector.
  /// - Returns: A dictionary.
  func toDictionary<K: Hashable>(keySelector: (Element) -> K) -> [K: Element] {
    var dict: [K: Element] = [:]
    for element in self {
      dict[keySelector(element)] = element
    }
    return dict
  }

  /// Transforms an array to a dictionary with each key-value pair being a key from the key selector and the value from the value selector. Inspired by RxJava's toMap.
  ///
  /// - Parameters:
  ///   - keySelector: Target key selector.
  ///   - valueSelector: Target value selector.
  /// - Returns: A dictionary.
  func toDictionary<K: Hashable, V: Any>(keySelector: (Element) -> K, valueSelector: (Element) -> V) -> [K: V] {
    var dict: [K: V] = [:]
    for element in self {
      dict[keySelector(element)] = valueSelector(element)
    }
    return dict
  }

  /// Transforms an array to a dictionary with each key-value pair being a key from the key selector and the value as an array of elements with the same key. Inspired by RxJava's toMultiMap.
  ///
  /// - Parameter keySelector: Target key selector.
  /// - Returns: A dictionary.
  func toMultiDictionary<K: Hashable>(keySelector: (Element) -> K) -> [K: [Element]] {
    var dict: [K: [Element]] = [:]
    for element in self {
      let key = keySelector(element)
      var values = dict[key]
      if values == nil {
        values = [element]
      } else {
        values?.append(element)
      }
      dict[key] = values!
    }
    return dict
  }

  /// Transforms an array to a dictionary with each key-value pair being a key from the key selector and the value as an array of value from the value selector with the same key. Inspired by RxJava's toMultiMap.
  ///
  /// - Parameters:
  ///   - keySelector: Target key selector.
  ///   - valueSelector: Target value selector.
  /// - Returns: A dictionary.
  func toMultiDictionary<K: Hashable, V: Any>(keySelector: (Element) -> K, valueSelector: (Element) -> V) -> [K: [V]] {
    var dict: [K: [V]] = [:]
    for element in self {
      let key = keySelector(element)
      let value = valueSelector(element)
      var values = dict[key]
      if values == nil {
        values = [value]
      } else {
        values?.append(value)
      }
      dict[key] = values!
    }
    return dict
  }
  
}
