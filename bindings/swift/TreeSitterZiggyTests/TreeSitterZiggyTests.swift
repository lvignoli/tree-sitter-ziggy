import XCTest
import SwiftTreeSitter
import TreeSitterZiggy

final class TreeSitterZiggyTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_ziggy())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Ziggy grammar")
    }
}
