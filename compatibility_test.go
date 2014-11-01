package xlsx

import . "gopkg.in/check.v1"

type GoogleDocsExcelSuite struct{}

var _ = Suite(&GoogleDocsExcelSuite{})

// Test that we can successfully read an XLSX file generated by
// Google Docs.
func (g *GoogleDocsExcelSuite) TestGoogleDocsExcel(c *C) {
	xlsxFile, err := OpenFile("./testdocs/googleDocsTest.xlsx")
	c.Assert(err, IsNil)
	c.Assert(xlsxFile, NotNil)
}

type MacExcelSuite struct{}

var _ = Suite(&MacExcelSuite{})

// Test that we can successfully read an XLSX file generated by
// Microsoft Excel for Mac.  In particular this requires that we
// respect the contents of workbook.xml.rels, which maps the sheet IDs
// to their internal file names.
func (m *MacExcelSuite) TestMacExcel(c *C) {
	xlsxFile, err := OpenFile("./testdocs/macExcelTest.xlsx")
	c.Assert(err, IsNil)
	c.Assert(xlsxFile, NotNil)
	s := xlsxFile.Sheet["普通技能"].Cell(0, 0).String()
	c.Assert(s, Equals, "编号")
}