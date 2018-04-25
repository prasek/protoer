package tests

import (
	"io/ioutil"
	"os"
	"testing"

	dpb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	_ "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
	_ "github.com/gogo/protobuf/types"
	"github.com/prasek/protoer/examples/desc"
	"github.com/prasek/protoer/internal/test/testutil"
	"github.com/prasek/protoer/proto"
)

func TestFileDescriptorObjectGraph(t *testing.T) {
	// This checks the structure of the descriptor for desc_test1.proto to make sure
	// the "rich descriptor" accurately models everything therein.
	fd, err := loadProtoset("../testprotos/desc_test1.protoset")
	testutil.Ok(t, err)
	checkDescriptor(t, "file", 0, fd, nil, fd, descCase{
		name: "desc_test1.proto",
		references: map[string]childCases{
			"messages": {(*desc.FileDescriptor).GetMessageTypes, []descCase{
				{
					name: "testprotos.TestMessage",
					references: map[string]childCases{
						"fields": {(*desc.MessageDescriptor).GetFields, []descCase{
							{
								name: "testprotos.TestMessage.nm",
								references: map[string]childCases{
									"message type": {(*desc.FieldDescriptor).GetMessageType, refs("testprotos.TestMessage.NestedMessage")},
									"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
									"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
								},
							},
							{
								name: "testprotos.TestMessage.anm",
								references: map[string]childCases{
									"message type": {(*desc.FieldDescriptor).GetMessageType, refs("testprotos.TestMessage.NestedMessage.AnotherNestedMessage")},
									"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
									"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
								},
							},
							{
								name: "testprotos.TestMessage.yanm",
								references: map[string]childCases{
									"message type": {(*desc.FieldDescriptor).GetMessageType, refs("testprotos.TestMessage.NestedMessage.AnotherNestedMessage.YetAnotherNestedMessage")},
									"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
									"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
								},
							},
							{
								name: "testprotos.TestMessage.ne",
								references: map[string]childCases{
									"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
									"enum type":    {(*desc.FieldDescriptor).GetEnumType, refs("testprotos.TestMessage.NestedEnum")},
									"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
								},
							},
						}},
						// this rabbit hole goes pretty deep...
						"nested messages": {(*desc.MessageDescriptor).GetNestedMessageTypes, []descCase{
							{
								name: "testprotos.TestMessage.NestedMessage",
								references: map[string]childCases{
									"fields": {(*desc.MessageDescriptor).GetFields, []descCase{
										{
											name: "testprotos.TestMessage.NestedMessage.anm",
											references: map[string]childCases{
												"message type": {(*desc.FieldDescriptor).GetMessageType, refs("testprotos.TestMessage.NestedMessage.AnotherNestedMessage")},
												"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
												"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
											},
										},
										{
											name: "testprotos.TestMessage.NestedMessage.yanm",
											references: map[string]childCases{
												"message type": {(*desc.FieldDescriptor).GetMessageType, refs("testprotos.TestMessage.NestedMessage.AnotherNestedMessage.YetAnotherNestedMessage")},
												"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
												"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
											},
										},
									}},
									"nested messages": {(*desc.MessageDescriptor).GetNestedMessageTypes, []descCase{
										{
											name: "testprotos.TestMessage.NestedMessage.AnotherNestedMessage",
											references: map[string]childCases{
												"fields": {(*desc.MessageDescriptor).GetFields, []descCase{
													{
														name: "testprotos.TestMessage.NestedMessage.AnotherNestedMessage.yanm",
														references: map[string]childCases{
															"message type": {(*desc.FieldDescriptor).GetMessageType, refs("testprotos.TestMessage.NestedMessage.AnotherNestedMessage.YetAnotherNestedMessage")},
															"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
															"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
														},
													},
												}},
												"nested messages": {(*desc.MessageDescriptor).GetNestedMessageTypes, []descCase{
													{
														name: "testprotos.TestMessage.NestedMessage.AnotherNestedMessage.YetAnotherNestedMessage",
														references: map[string]childCases{
															"nested fields": {(*desc.MessageDescriptor).GetFields, []descCase{
																{
																	name: "testprotos.TestMessage.NestedMessage.AnotherNestedMessage.YetAnotherNestedMessage.foo",
																	references: map[string]childCases{
																		"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
																		"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
																		"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
																	},
																},
																{
																	name: "testprotos.TestMessage.NestedMessage.AnotherNestedMessage.YetAnotherNestedMessage.bar",
																	references: map[string]childCases{
																		"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
																		"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
																		"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
																	},
																},
																{
																	name: "testprotos.TestMessage.NestedMessage.AnotherNestedMessage.YetAnotherNestedMessage.baz",
																	references: map[string]childCases{
																		"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
																		"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
																		"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
																	},
																},
																{
																	name: "testprotos.TestMessage.NestedMessage.AnotherNestedMessage.YetAnotherNestedMessage.dne",
																	references: map[string]childCases{
																		"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
																		"enum type":    {(*desc.FieldDescriptor).GetEnumType, refs("testprotos.TestMessage.NestedMessage.AnotherNestedMessage.YetAnotherNestedMessage.DeeplyNestedEnum")},
																		"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
																	},
																},
																{
																	name: "testprotos.TestMessage.NestedMessage.AnotherNestedMessage.YetAnotherNestedMessage.anm",
																	references: map[string]childCases{
																		"message type": {(*desc.FieldDescriptor).GetMessageType, refs("testprotos.TestMessage.NestedMessage.AnotherNestedMessage")},
																		"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
																		"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
																	},
																},
																{
																	name: "testprotos.TestMessage.NestedMessage.AnotherNestedMessage.YetAnotherNestedMessage.nm",
																	references: map[string]childCases{
																		"message type": {(*desc.FieldDescriptor).GetMessageType, refs("testprotos.TestMessage.NestedMessage")},
																		"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
																		"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
																	},
																},
																{
																	name: "testprotos.TestMessage.NestedMessage.AnotherNestedMessage.YetAnotherNestedMessage.tm",
																	references: map[string]childCases{
																		"message type": {(*desc.FieldDescriptor).GetMessageType, refs("testprotos.TestMessage")},
																		"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
																		"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
																	},
																},
															}},
															"nested messages": {(*desc.MessageDescriptor).GetNestedMessageTypes, nil},
															"nested enums": {(*desc.MessageDescriptor).GetNestedEnumTypes, []descCase{
																{
																	name: "testprotos.TestMessage.NestedMessage.AnotherNestedMessage.YetAnotherNestedMessage.DeeplyNestedEnum",
																	references: map[string]childCases{
																		"values": {(*desc.EnumDescriptor).GetValues, children(
																			"testprotos.TestMessage.NestedMessage.AnotherNestedMessage.YetAnotherNestedMessage.DeeplyNestedEnum.VALUE1",
																			"testprotos.TestMessage.NestedMessage.AnotherNestedMessage.YetAnotherNestedMessage.DeeplyNestedEnum.VALUE2"),
																		},
																	},
																},
															}},
															"nested extensions": {(*desc.MessageDescriptor).GetNestedExtensions, nil},
															"one ofs":           {(*desc.MessageDescriptor).GetOneOfs, nil},
														},
													},
												}},
												"nested enums": {(*desc.MessageDescriptor).GetNestedEnumTypes, nil},
												"nested extensions": {(*desc.MessageDescriptor).GetNestedExtensions, []descCase{
													{
														name:   "testprotos.TestMessage.NestedMessage.AnotherNestedMessage.flags",
														number: 200,
														references: map[string]childCases{
															"owner":        {(*desc.FieldDescriptor).GetOwner, refs("testprotos.AnotherTestMessage")},
															"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
															"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
															"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
														},
													},
												}},
												"one ofs": {(*desc.MessageDescriptor).GetOneOfs, nil},
											},
										},
									}},
									"nested enums":      {(*desc.MessageDescriptor).GetNestedEnumTypes, nil},
									"nested extensions": {(*desc.MessageDescriptor).GetNestedExtensions, nil},
									"one ofs":           {(*desc.MessageDescriptor).GetOneOfs, nil},
								},
							},
						}},
						"nested enums": {(*desc.MessageDescriptor).GetNestedEnumTypes, []descCase{
							{
								name: "testprotos.TestMessage.NestedEnum",
								references: map[string]childCases{
									"values": {(*desc.EnumDescriptor).GetValues, children(
										"testprotos.TestMessage.NestedEnum.VALUE1", "testprotos.TestMessage.NestedEnum.VALUE2"),
									},
								},
							},
						}},
						"nested extensions": {(*desc.MessageDescriptor).GetNestedExtensions, nil},
						"one ofs":           {(*desc.MessageDescriptor).GetOneOfs, nil},
					},
				},
				{
					name: "testprotos.AnotherTestMessage",
					references: map[string]childCases{
						"fields": {(*desc.MessageDescriptor).GetFields, []descCase{
							{
								name: "testprotos.AnotherTestMessage.dne",
								references: map[string]childCases{
									"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
									"enum type":    {(*desc.FieldDescriptor).GetEnumType, refs("testprotos.TestMessage.NestedMessage.AnotherNestedMessage.YetAnotherNestedMessage.DeeplyNestedEnum")},
									"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
								},
							},
							{
								name: "testprotos.AnotherTestMessage.map_field1",
								references: map[string]childCases{
									"message type": {(*desc.FieldDescriptor).GetMessageType, refs("testprotos.AnotherTestMessage.MapField1Entry")},
									"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
									"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
								},
							},
							{
								name: "testprotos.AnotherTestMessage.map_field2",
								references: map[string]childCases{
									"message type": {(*desc.FieldDescriptor).GetMessageType, refs("testprotos.AnotherTestMessage.MapField2Entry")},
									"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
									"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
								},
							},
							{
								name: "testprotos.AnotherTestMessage.map_field3",
								references: map[string]childCases{
									"message type": {(*desc.FieldDescriptor).GetMessageType, refs("testprotos.AnotherTestMessage.MapField3Entry")},
									"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
									"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
								},
							},
							{
								name: "testprotos.AnotherTestMessage.map_field4",
								references: map[string]childCases{
									"message type": {(*desc.FieldDescriptor).GetMessageType, refs("testprotos.AnotherTestMessage.MapField4Entry")},
									"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
									"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
								},
							},
							{
								name: "testprotos.AnotherTestMessage.rocknroll",
								references: map[string]childCases{
									"message type": {(*desc.FieldDescriptor).GetMessageType, refs("testprotos.AnotherTestMessage.RockNRoll")},
									"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
									"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
								},
							},
							{
								name: "testprotos.AnotherTestMessage.str",
								references: map[string]childCases{
									"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
									"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
									"one of":       {(*desc.FieldDescriptor).GetOneOf, refs("testprotos.AnotherTestMessage.atmoo")},
								},
							},
							{
								name: "testprotos.AnotherTestMessage.int",
								references: map[string]childCases{
									"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
									"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
									"one of":       {(*desc.FieldDescriptor).GetOneOf, refs("testprotos.AnotherTestMessage.atmoo")},
								},
							},
						}},
						"one ofs": {(*desc.MessageDescriptor).GetOneOfs, []descCase{
							{
								name:       "testprotos.AnotherTestMessage.atmoo",
								skipParent: true,
								references: map[string]childCases{
									"fields": {(*desc.OneOfDescriptor).GetChoices, fields(
										fld{"testprotos.AnotherTestMessage.str", 7},
										fld{"testprotos.AnotherTestMessage.int", 8}),
									},
								},
							},
						}},
						"nested messages": {(*desc.MessageDescriptor).GetNestedMessageTypes, []descCase{
							{
								name: "testprotos.AnotherTestMessage.MapField1Entry",
								references: map[string]childCases{
									"fields": {(*desc.MessageDescriptor).GetFields, []descCase{
										{
											name: "testprotos.AnotherTestMessage.MapField1Entry.key",
											references: map[string]childCases{
												"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
												"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
												"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
											},
										},
										{
											name: "testprotos.AnotherTestMessage.MapField1Entry.value",
											references: map[string]childCases{
												"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
												"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
												"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
											},
										},
									}},
									"nested messages":   {(*desc.MessageDescriptor).GetNestedMessageTypes, nil},
									"nested enums":      {(*desc.MessageDescriptor).GetNestedEnumTypes, nil},
									"nested extensions": {(*desc.MessageDescriptor).GetNestedExtensions, nil},
									"one ofs":           {(*desc.MessageDescriptor).GetOneOfs, nil},
								},
							},
							{
								name: "testprotos.AnotherTestMessage.MapField2Entry",
								references: map[string]childCases{
									"fields": {(*desc.MessageDescriptor).GetFields, []descCase{
										{
											name: "testprotos.AnotherTestMessage.MapField2Entry.key",
											references: map[string]childCases{
												"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
												"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
												"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
											},
										},
										{
											name: "testprotos.AnotherTestMessage.MapField2Entry.value",
											references: map[string]childCases{
												"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
												"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
												"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
											},
										},
									}},
									"nested messages":   {(*desc.MessageDescriptor).GetNestedMessageTypes, nil},
									"nested enums":      {(*desc.MessageDescriptor).GetNestedEnumTypes, nil},
									"nested extensions": {(*desc.MessageDescriptor).GetNestedExtensions, nil},
									"one ofs":           {(*desc.MessageDescriptor).GetOneOfs, nil},
								},
							},
							{
								name: "testprotos.AnotherTestMessage.MapField3Entry",
								references: map[string]childCases{
									"fields": {(*desc.MessageDescriptor).GetFields, []descCase{
										{
											name: "testprotos.AnotherTestMessage.MapField3Entry.key",
											references: map[string]childCases{
												"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
												"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
												"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
											},
										},
										{
											name: "testprotos.AnotherTestMessage.MapField3Entry.value",
											references: map[string]childCases{
												"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
												"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
												"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
											},
										},
									}},
									"nested messages":   {(*desc.MessageDescriptor).GetNestedMessageTypes, nil},
									"nested enums":      {(*desc.MessageDescriptor).GetNestedEnumTypes, nil},
									"nested extensions": {(*desc.MessageDescriptor).GetNestedExtensions, nil},
									"one ofs":           {(*desc.MessageDescriptor).GetOneOfs, nil},
								},
							},
							{
								name: "testprotos.AnotherTestMessage.MapField4Entry",
								references: map[string]childCases{
									"fields": {(*desc.MessageDescriptor).GetFields, []descCase{
										{
											name: "testprotos.AnotherTestMessage.MapField4Entry.key",
											references: map[string]childCases{
												"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
												"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
												"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
											},
										},
										{
											name: "testprotos.AnotherTestMessage.MapField4Entry.value",
											references: map[string]childCases{
												"message type": {(*desc.FieldDescriptor).GetMessageType, refs("testprotos.AnotherTestMessage")},
												"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
												"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
											},
										},
									}},
									"nested messages":   {(*desc.MessageDescriptor).GetNestedMessageTypes, nil},
									"nested enums":      {(*desc.MessageDescriptor).GetNestedEnumTypes, nil},
									"nested extensions": {(*desc.MessageDescriptor).GetNestedExtensions, nil},
									"one ofs":           {(*desc.MessageDescriptor).GetOneOfs, nil},
								},
							},
							{
								name: "testprotos.AnotherTestMessage.RockNRoll",
								references: map[string]childCases{
									"fields": {(*desc.MessageDescriptor).GetFields, []descCase{
										{
											name: "testprotos.AnotherTestMessage.RockNRoll.beatles",
											references: map[string]childCases{
												"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
												"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
												"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
											},
										},
										{
											name: "testprotos.AnotherTestMessage.RockNRoll.stones",
											references: map[string]childCases{
												"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
												"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
												"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
											},
										},
										{
											name: "testprotos.AnotherTestMessage.RockNRoll.doors",
											references: map[string]childCases{
												"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
												"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
												"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
											},
										},
									}},
									"nested messages":   {(*desc.MessageDescriptor).GetNestedMessageTypes, nil},
									"nested enums":      {(*desc.MessageDescriptor).GetNestedEnumTypes, nil},
									"nested extensions": {(*desc.MessageDescriptor).GetNestedExtensions, nil},
									"one ofs":           {(*desc.MessageDescriptor).GetOneOfs, nil},
								},
							},
						}},
						"nested enums":      {(*desc.MessageDescriptor).GetNestedEnumTypes, nil},
						"nested extensions": {(*desc.MessageDescriptor).GetNestedExtensions, nil},
					},
				},
			}},
			"enums":    {(*desc.FileDescriptor).GetEnumTypes, nil},
			"services": {(*desc.FileDescriptor).GetServices, nil},
			"extensions": {(*desc.FileDescriptor).GetExtensions, []descCase{
				{
					name:   "testprotos.xtm",
					number: 100,
					references: map[string]childCases{
						"owner":        {(*desc.FieldDescriptor).GetOwner, refs("testprotos.AnotherTestMessage")},
						"message type": {(*desc.FieldDescriptor).GetMessageType, refs("testprotos.TestMessage")},
						"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
						"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
					},
				},
				{
					name:   "testprotos.xs",
					number: 101,
					references: map[string]childCases{
						"owner":        {(*desc.FieldDescriptor).GetOwner, refs("testprotos.AnotherTestMessage")},
						"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
						"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
						"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
					},
				},
				{
					name:   "testprotos.xi",
					number: 102,
					references: map[string]childCases{
						"owner":        {(*desc.FieldDescriptor).GetOwner, refs("testprotos.AnotherTestMessage")},
						"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
						"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
						"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
					},
				},
				{
					name:   "testprotos.xui",
					number: 103,
					references: map[string]childCases{
						"owner":        {(*desc.FieldDescriptor).GetOwner, refs("testprotos.AnotherTestMessage")},
						"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
						"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
						"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
					},
				},
			}},
		},
	})
}

func TestOneOfDescriptors(t *testing.T) {
	fd, err := desc.LoadFileDescriptor("desc_test2.proto")
	testutil.Ok(t, err)
	md, err := desc.LoadMessageDescriptor("testprotos.Frobnitz")
	testutil.Ok(t, err)
	checkDescriptor(t, "message", 0, md, fd, fd, descCase{
		name: "testprotos.Frobnitz",
		references: map[string]childCases{
			"fields": {(*desc.MessageDescriptor).GetFields, []descCase{
				{
					name: "testprotos.Frobnitz.a",
					references: map[string]childCases{
						"message type": {(*desc.FieldDescriptor).GetMessageType, refs("testprotos.TestMessage")},
						"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
						"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
					},
				},
				{
					name: "testprotos.Frobnitz.b",
					references: map[string]childCases{
						"message type": {(*desc.FieldDescriptor).GetMessageType, refs("testprotos.AnotherTestMessage")},
						"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
						"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
					},
				},
				{
					name: "testprotos.Frobnitz.c1",
					references: map[string]childCases{
						"message type": {(*desc.FieldDescriptor).GetMessageType, refs("testprotos.TestMessage.NestedMessage")},
						"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
						"one of":       {(*desc.FieldDescriptor).GetOneOf, refs("testprotos.Frobnitz.abc")},
					},
				},
				{
					name: "testprotos.Frobnitz.c2",
					references: map[string]childCases{
						"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
						"enum type":    {(*desc.FieldDescriptor).GetEnumType, refs("testprotos.TestMessage.NestedEnum")},
						"one of":       {(*desc.FieldDescriptor).GetOneOf, refs("testprotos.Frobnitz.abc")},
					},
				},
				{
					name: "testprotos.Frobnitz.d",
					references: map[string]childCases{
						"message type": {(*desc.FieldDescriptor).GetMessageType, refs("testprotos.TestMessage.NestedMessage")},
						"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
						"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
					},
				},
				{
					name: "testprotos.Frobnitz.e",
					references: map[string]childCases{
						"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
						"enum type":    {(*desc.FieldDescriptor).GetEnumType, refs("testprotos.TestMessage.NestedEnum")},
						"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
					},
				},
				{
					name: "testprotos.Frobnitz.f",
					references: map[string]childCases{
						"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
						"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
						"one of":       {(*desc.FieldDescriptor).GetOneOf, nil},
					},
				},
				{
					name: "testprotos.Frobnitz.g1",
					references: map[string]childCases{
						"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
						"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
						"one of":       {(*desc.FieldDescriptor).GetOneOf, refs("testprotos.Frobnitz.def")},
					},
				},
				{
					name: "testprotos.Frobnitz.g2",
					references: map[string]childCases{
						"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
						"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
						"one of":       {(*desc.FieldDescriptor).GetOneOf, refs("testprotos.Frobnitz.def")},
					},
				},
				{
					name: "testprotos.Frobnitz.g3",
					references: map[string]childCases{
						"message type": {(*desc.FieldDescriptor).GetMessageType, nil},
						"enum type":    {(*desc.FieldDescriptor).GetEnumType, nil},
						"one of":       {(*desc.FieldDescriptor).GetOneOf, refs("testprotos.Frobnitz.def")},
					},
				},
			}},
			"nested messages":   {(*desc.MessageDescriptor).GetNestedMessageTypes, nil},
			"nested enums":      {(*desc.MessageDescriptor).GetNestedEnumTypes, nil},
			"nested extensions": {(*desc.MessageDescriptor).GetNestedExtensions, nil},
			"one ofs": {(*desc.MessageDescriptor).GetOneOfs, []descCase{
				{
					name:       "testprotos.Frobnitz.abc",
					skipParent: true,
					references: map[string]childCases{
						"fields": {(*desc.OneOfDescriptor).GetChoices, fields(
							fld{"testprotos.Frobnitz.c1", 3},
							fld{"testprotos.Frobnitz.c2", 4}),
						},
					},
				},
				{
					name:       "testprotos.Frobnitz.def",
					skipParent: true,
					references: map[string]childCases{
						"fields": {(*desc.OneOfDescriptor).GetChoices, fields(
							fld{"testprotos.Frobnitz.g1", 8},
							fld{"testprotos.Frobnitz.g2", 9},
							fld{"testprotos.Frobnitz.g3", 10}),
						},
					},
				},
			}},
		},
	})
}

func TestLoadFileDescriptorWithDeps(t *testing.T) {
	// Try one with some imports
	fd, err := desc.LoadFileDescriptor("desc_test2.proto")
	testutil.Ok(t, err)
	testutil.Eq(t, "desc_test2.proto", fd.GetName())
	testutil.Eq(t, "desc_test2.proto", fd.GetFullyQualifiedName())
	testutil.Eq(t, "testprotos", fd.GetPackage())

	deps := fd.GetDependencies()
	testutil.Eq(t, 3, len(deps))
	testutil.Eq(t, "desc_test1.proto", deps[0].GetName())
	testutil.Eq(t, "pkg/desc_test_pkg.proto", deps[1].GetName())
	testutil.Eq(t, "nopkg/desc_test_nopkg.proto", deps[2].GetName())

	// loading the dependencies yields same descriptor objects
	fd, err = desc.LoadFileDescriptor("desc_test1.proto")
	testutil.Ok(t, err)
	testutil.Eq(t, deps[0], fd)
	fd, err = desc.LoadFileDescriptor("pkg/desc_test_pkg.proto")
	testutil.Ok(t, err)
	testutil.Eq(t, deps[1], fd)
	fd, err = desc.LoadFileDescriptor("nopkg/desc_test_nopkg.proto")
	testutil.Ok(t, err)
	testutil.Eq(t, deps[2], fd)
}

func TestLoadFileDescriptorForWellKnownProtos(t *testing.T) {
	wellKnownProtos := map[string][]string{
		"google/protobuf/any.proto":             {"google.protobuf.Any"},
		"google/protobuf/api.proto":             {"google.protobuf.Api", "google.protobuf.Method", "google.protobuf.Mixin"},
		"google/protobuf/descriptor.proto":      {"google.protobuf.FileDescriptorSet", "google.protobuf.DescriptorProto"},
		"google/protobuf/duration.proto":        {"google.protobuf.Duration"},
		"google/protobuf/empty.proto":           {"google.protobuf.Empty"},
		"google/protobuf/field_mask.proto":      {"google.protobuf.FieldMask"},
		"google/protobuf/source_context.proto":  {"google.protobuf.SourceContext"},
		"google/protobuf/struct.proto":          {"google.protobuf.Struct", "google.protobuf.Value", "google.protobuf.NullValue"},
		"google/protobuf/timestamp.proto":       {"google.protobuf.Timestamp"},
		"google/protobuf/type.proto":            {"google.protobuf.Type", "google.protobuf.Field", "google.protobuf.Syntax"},
		"google/protobuf/wrappers.proto":        {"google.protobuf.DoubleValue", "google.protobuf.Int32Value", "google.protobuf.StringValue"},
		"google/protobuf/compiler/plugin.proto": {"google.protobuf.compiler.CodeGeneratorRequest"},
	}

	aliases := proto.Aliases()

	for file, types := range wellKnownProtos {
		fd, err := desc.LoadFileDescriptor(file)
		testutil.Ok(t, err)
		testutil.Eq(t, file, fd.GetName())
		for _, typ := range types {
			d := fd.FindSymbol(typ)
			testutil.Require(t, d != nil)
		}

		// also try loading via alternate name
		if aliases == nil {
			continue
		}
		file = aliases[file]
		if file == "" {
			// not a file that has a known alternate, so nothing else to check...
			continue
		}
		fd, err = desc.LoadFileDescriptor(file)
		testutil.Ok(t, err)
		testutil.Eq(t, file, fd.GetName())
		for _, typ := range types {
			d := fd.FindSymbol(typ)
			testutil.Require(t, d != nil)
		}
	}
}

func TestFileDescriptorDeps(t *testing.T) {
	// tests accessors for public and weak dependencies
	fd1 := createDesc(t, &dpb.FileDescriptorProto{Name: proto.String("a")})
	fd2 := createDesc(t, &dpb.FileDescriptorProto{Name: proto.String("b")})
	fd3 := createDesc(t, &dpb.FileDescriptorProto{Name: proto.String("c")})
	fd4 := createDesc(t, &dpb.FileDescriptorProto{Name: proto.String("d")})
	fd5 := createDesc(t, &dpb.FileDescriptorProto{Name: proto.String("e")})
	fd := createDesc(t, &dpb.FileDescriptorProto{
		Name:             proto.String("f"),
		Dependency:       []string{"a", "b", "c", "d", "e"},
		PublicDependency: []int32{1, 3},
		WeakDependency:   []int32{2, 4},
	}, fd1, fd2, fd3, fd4, fd5)

	deps := fd.GetDependencies()
	testutil.Eq(t, 5, len(deps))
	testutil.Eq(t, fd1, deps[0])
	testutil.Eq(t, fd2, deps[1])
	testutil.Eq(t, fd3, deps[2])
	testutil.Eq(t, fd4, deps[3])
	testutil.Eq(t, fd5, deps[4])

	deps = fd.GetPublicDependencies()
	testutil.Eq(t, 2, len(deps))
	testutil.Eq(t, fd2, deps[0])
	testutil.Eq(t, fd4, deps[1])

	deps = fd.GetWeakDependencies()
	testutil.Eq(t, 2, len(deps))
	testutil.Eq(t, fd3, deps[0])
	testutil.Eq(t, fd5, deps[1])

	// Now try on a simple descriptor emitted by protoc
	fd6, err := desc.LoadFileDescriptor("nopkg/desc_test_nopkg.proto")
	testutil.Ok(t, err)
	fd7, err := desc.LoadFileDescriptor("nopkg/desc_test_nopkg_new.proto")
	testutil.Ok(t, err)
	deps = fd6.GetPublicDependencies()
	testutil.Eq(t, 1, len(deps))
	testutil.Eq(t, fd7, deps[0])
}

func createDesc(t *testing.T, fd *dpb.FileDescriptorProto, deps ...*desc.FileDescriptor) *desc.FileDescriptor {
	desc, err := desc.CreateFileDescriptor(fd, deps...)
	testutil.Ok(t, err)
	return desc
}

func TestLoadFileDescriptor(t *testing.T) {
	fd, err := desc.LoadFileDescriptor("desc_test1.proto")
	testutil.Ok(t, err)
	// some very shallow tests (we have more detailed ones in other test cases)
	testutil.Eq(t, "desc_test1.proto", fd.GetName())
	testutil.Eq(t, "desc_test1.proto", fd.GetFullyQualifiedName())
	testutil.Eq(t, "testprotos", fd.GetPackage())
}

func loadProtoset(path string) (*desc.FileDescriptor, error) {
	var fds dpb.FileDescriptorSet
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	bb, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	if err = proto.Unmarshal(bb, &fds); err != nil {
		return nil, err
	}
	return desc.CreateFileDescriptorFromSet(&fds)
}
