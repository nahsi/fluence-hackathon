// Copyright 2023 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package index

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/pb33f/libopenapi/utils"
	"gopkg.in/yaml.v3"
)

func (index *SpecIndex) extractDefinitionsAndSchemas(schemasNode *yaml.Node, pathPrefix string) {
	var name string
	for i, schema := range schemasNode.Content {
		if i%2 == 0 {
			name = schema.Value
			continue
		}

		def := fmt.Sprintf("%s%s", pathPrefix, name)
		fullDef := fmt.Sprintf("%s%s", index.specAbsolutePath, def)

		ref := &Reference{
			FullDefinition:        fullDef,
			Definition:            def,
			Name:                  name,
			Node:                  schema,
			Path:                  fmt.Sprintf("$.components.schemas['%s']", name),
			ParentNode:            schemasNode,
			RequiredRefProperties: extractDefinitionRequiredRefProperties(schemasNode, map[string][]string{}, fullDef, index),
		}
		index.allComponentSchemaDefinitions.Store(def, ref)
	}
}

// extractDefinitionRequiredRefProperties goes through the direct properties of a schema and extracts the map of required definitions from within it
func extractDefinitionRequiredRefProperties(schemaNode *yaml.Node, reqRefProps map[string][]string, fulldef string, idx *SpecIndex) map[string][]string {
	if schemaNode == nil {
		return reqRefProps
	}

	// If the node we're looking at is a direct ref to another model without any properties, mark it as required, but still continue to look for required properties
	isRef, _, defPath := utils.IsNodeRefValue(schemaNode)
	if isRef {
		if _, ok := reqRefProps[defPath]; !ok {
			reqRefProps[defPath] = []string{}
		}
	}

	// Check for a required parameters list, and return if none exists, as any properties will be optional
	_, requiredSeqNode := utils.FindKeyNodeTop("required", schemaNode.Content)
	if requiredSeqNode == nil {
		return reqRefProps
	}

	_, propertiesMapNode := utils.FindKeyNodeTop("properties", schemaNode.Content)
	if propertiesMapNode == nil {
		// TODO: Log a warning on the resolver, because if you have required properties, but no actual properties, something is wrong
		return reqRefProps
	}

	name := ""
	for i, param := range propertiesMapNode.Content {
		if i%2 == 0 {
			name = param.Value
			continue
		}

		// Check to see if the current property is directly embedded within the current schema, and handle its properties if so
		_, paramPropertiesMapNode := utils.FindKeyNodeTop("properties", param.Content)
		if paramPropertiesMapNode != nil {
			reqRefProps = extractDefinitionRequiredRefProperties(param, reqRefProps, fulldef, idx)
		}

		// Check to see if the current property is polymorphic, and dive into that model if so
		for _, key := range []string{"allOf", "oneOf", "anyOf"} {
			_, ofNode := utils.FindKeyNodeTop(key, param.Content)
			if ofNode != nil {
				for _, ofNodeItem := range ofNode.Content {
					reqRefProps = extractRequiredReferenceProperties(fulldef, idx, ofNodeItem, name, reqRefProps)
				}
			}
		}
	}

	// Run through each of the required properties and extract _their_ required references
	for _, requiredPropertyNode := range requiredSeqNode.Content {
		_, requiredPropDefNode := utils.FindKeyNodeTop(requiredPropertyNode.Value, propertiesMapNode.Content)
		if requiredPropDefNode == nil {
			continue
		}

		reqRefProps = extractRequiredReferenceProperties(fulldef, idx, requiredPropDefNode, requiredPropertyNode.Value, reqRefProps)
	}

	return reqRefProps
}

// extractRequiredReferenceProperties returns a map of definition names to the property or properties which reference it within a node
func extractRequiredReferenceProperties(fulldef string, idx *SpecIndex, requiredPropDefNode *yaml.Node, propName string, reqRefProps map[string][]string) map[string][]string {
	isRef, _, refName := utils.IsNodeRefValue(requiredPropDefNode)
	if !isRef {
		_, defItems := utils.FindKeyNodeTop("items", requiredPropDefNode.Content)
		if defItems != nil {
			isRef, _, refName = utils.IsNodeRefValue(defItems)
		}
	}

	if /* still */ !isRef {
		return reqRefProps
	}

	defPath := fulldef

	if strings.HasPrefix(refName, "http") || filepath.IsAbs(refName) {
		defPath = refName
	} else {
		exp := strings.Split(fulldef, "#/")
		if len(exp) == 2 {
			if exp[0] != "" {
				if strings.HasPrefix(exp[0], "http") {
					u, _ := url.Parse(exp[0])
					r := strings.Split(refName, "#/")
					if len(r) == 2 {
						var abs string
						if r[0] == "" {
							abs = u.Path
						} else {
							abs, _ = filepath.Abs(utils.CheckPathOverlap(filepath.Dir(u.Path), r[0],
								string(os.PathSeparator)))
						}

						u.Path = utils.ReplaceWindowsDriveWithLinuxPath(abs)
						u.Fragment = ""
						defPath = fmt.Sprintf("%s#/%s", u.String(), r[1])
					} else {
						u.Path = utils.ReplaceWindowsDriveWithLinuxPath(utils.CheckPathOverlap(filepath.Dir(u.Path),
							r[0], string(os.PathSeparator)))
						u.Fragment = ""
						defPath = u.String()
					}
				} else {
					r := strings.Split(refName, "#/")
					if len(r) == 2 {
						var abs string
						if r[0] == "" {
							abs, _ = filepath.Abs(exp[0])
						} else {
							abs, _ = filepath.Abs(utils.CheckPathOverlap(filepath.Dir(exp[0]), r[0],
								string(os.PathSeparator)))

							//abs, _ = filepath.Abs(filepath.Join(filepath.Dir(exp[0]), r[0],
							//	string('J')))
						}

						defPath = fmt.Sprintf("%s#/%s", abs, r[1])
					} else {
						defPath, _ = filepath.Abs(utils.CheckPathOverlap(filepath.Dir(exp[0]),
							r[0], string(os.PathSeparator)))
					}
				}
			} else {
				defPath = refName
			}
		} else {
			if strings.HasPrefix(exp[0], "http") {
				u, _ := url.Parse(exp[0])
				r := strings.Split(refName, "#/")
				if len(r) == 2 {
					abs, _ := filepath.Abs(utils.CheckPathOverlap(filepath.Dir(u.Path), r[0], string(os.PathSeparator)))
					u.Path = utils.ReplaceWindowsDriveWithLinuxPath(abs)
					u.Fragment = ""
					defPath = fmt.Sprintf("%s#/%s", u.String(), r[1])
				} else {
					u.Path = utils.ReplaceWindowsDriveWithLinuxPath(utils.CheckPathOverlap(filepath.Dir(u.Path),
						r[0], string(os.PathSeparator)))
					u.Fragment = ""
					defPath = u.String()
				}
			} else {
				defPath, _ = filepath.Abs(utils.CheckPathOverlap(filepath.Dir(exp[0]), refName, string(os.PathSeparator)))
			}
		}
	}

	if _, ok := reqRefProps[defPath]; !ok {
		reqRefProps[defPath] = []string{}
	}
	reqRefProps[defPath] = append(reqRefProps[defPath], propName)

	return reqRefProps
}

func (index *SpecIndex) extractComponentParameters(paramsNode *yaml.Node, pathPrefix string) {
	var name string
	for i, param := range paramsNode.Content {
		if i%2 == 0 {
			name = param.Value
			continue
		}
		def := fmt.Sprintf("%s%s", pathPrefix, name)
		ref := &Reference{
			Definition: def,
			Name:       name,
			Node:       param,
		}
		index.allParameters[def] = ref
	}
}

func (index *SpecIndex) extractComponentRequestBodies(requestBodiesNode *yaml.Node, pathPrefix string) {
	var name string
	for i, reqBod := range requestBodiesNode.Content {
		if i%2 == 0 {
			name = reqBod.Value
			continue
		}
		def := fmt.Sprintf("%s%s", pathPrefix, name)
		ref := &Reference{
			Definition: def,
			Name:       name,
			Node:       reqBod,
		}
		index.allRequestBodies[def] = ref
	}
}

func (index *SpecIndex) extractComponentResponses(responsesNode *yaml.Node, pathPrefix string) {
	var name string
	for i, response := range responsesNode.Content {
		if i%2 == 0 {
			name = response.Value
			continue
		}
		def := fmt.Sprintf("%s%s", pathPrefix, name)
		ref := &Reference{
			Definition: def,
			Name:       name,
			Node:       response,
		}
		index.allResponses[def] = ref
	}
}

func (index *SpecIndex) extractComponentHeaders(headersNode *yaml.Node, pathPrefix string) {
	var name string
	for i, header := range headersNode.Content {
		if i%2 == 0 {
			name = header.Value
			continue
		}
		def := fmt.Sprintf("%s%s", pathPrefix, name)
		ref := &Reference{
			Definition: def,
			Name:       name,
			Node:       header,
		}
		index.allHeaders[def] = ref
	}
}

func (index *SpecIndex) extractComponentCallbacks(callbacksNode *yaml.Node, pathPrefix string) {
	var name string
	for i, callback := range callbacksNode.Content {
		if i%2 == 0 {
			name = callback.Value
			continue
		}
		def := fmt.Sprintf("%s%s", pathPrefix, name)
		ref := &Reference{
			Definition: def,
			Name:       name,
			Node:       callback,
		}
		index.allCallbacks[def] = ref
	}
}

func (index *SpecIndex) extractComponentLinks(linksNode *yaml.Node, pathPrefix string) {
	var name string
	for i, link := range linksNode.Content {
		if i%2 == 0 {
			name = link.Value
			continue
		}
		def := fmt.Sprintf("%s%s", pathPrefix, name)
		ref := &Reference{
			Definition: def,
			Name:       name,
			Node:       link,
		}
		index.allLinks[def] = ref
	}
}

func (index *SpecIndex) extractComponentExamples(examplesNode *yaml.Node, pathPrefix string) {
	var name string
	for i, example := range examplesNode.Content {
		if i%2 == 0 {
			name = example.Value
			continue
		}
		def := fmt.Sprintf("%s%s", pathPrefix, name)
		ref := &Reference{
			Definition: def,
			Name:       name,
			Node:       example,
		}
		index.allExamples[def] = ref
	}
}

func (index *SpecIndex) extractComponentSecuritySchemes(securitySchemesNode *yaml.Node, pathPrefix string) {
	var name string
	for i, schema := range securitySchemesNode.Content {
		if i%2 == 0 {
			name = schema.Value
			continue
		}
		def := fmt.Sprintf("%s%s", pathPrefix, name)
		fullDef := fmt.Sprintf("%s%s", index.specAbsolutePath, def)

		ref := &Reference{
			FullDefinition:        fullDef,
			Definition:            def,
			Name:                  name,
			Node:                  schema,
			Path:                  fmt.Sprintf("$.components.securitySchemes.%s", name),
			ParentNode:            securitySchemesNode,
			RequiredRefProperties: extractDefinitionRequiredRefProperties(securitySchemesNode, map[string][]string{}, fullDef, index),
		}
		index.allSecuritySchemes[def] = ref
	}
}

func (index *SpecIndex) countUniqueInlineDuplicates() int {
	if index.componentsInlineParamUniqueCount > 0 {
		return index.componentsInlineParamUniqueCount
	}
	unique := 0
	for _, p := range index.paramInlineDuplicateNames {
		if len(p) == 1 {
			unique++
		}
	}
	index.componentsInlineParamUniqueCount = unique
	return unique
}

func seekRefEnd(index *SpecIndex, refName string) *Reference {
	ref, _ := index.SearchIndexForReference(refName)
	if ref != nil {
		if ok, _, v := utils.IsNodeRefValue(ref.Node); ok {
			return seekRefEnd(ref.Index, v)
		}
	}
	return ref
}

func (index *SpecIndex) scanOperationParams(params []*yaml.Node, pathItemNode *yaml.Node, method string) {
	for i, param := range params {
		// param is ref
		if len(param.Content) > 0 && param.Content[0].Value == "$ref" {

			paramRefName := param.Content[1].Value
			paramRef := index.allMappedRefs[paramRefName]
			if paramRef == nil {
				// could be in the rolodex
				ref := seekRefEnd(index, paramRefName)
				if ref != nil {
					paramRef = ref
					if strings.Contains(paramRefName, "%") {
						paramRefName, _ = url.QueryUnescape(paramRefName)
					}
				}
			}

			if index.paramOpRefs[pathItemNode.Value] == nil {
				index.paramOpRefs[pathItemNode.Value] = make(map[string]map[string][]*Reference)
				index.paramOpRefs[pathItemNode.Value][method] = make(map[string][]*Reference)

			}
			// if we know the path, but it's a new method
			if index.paramOpRefs[pathItemNode.Value][method] == nil {
				index.paramOpRefs[pathItemNode.Value][method] = make(map[string][]*Reference)
			}

			// if this is a duplicate, add an error and ignore it
			if index.paramOpRefs[pathItemNode.Value][method][paramRefName] != nil {
				path := fmt.Sprintf("$.paths['%s'].%s.parameters[%d]", pathItemNode.Value, method, i)
				if method == "top" {
					path = fmt.Sprintf("$.paths['%s'].parameters[%d]", pathItemNode.Value, i)
				}

				index.operationParamErrors = append(index.operationParamErrors, &IndexingError{
					Err: fmt.Errorf("the `%s` operation parameter at path `%s`, "+
						"index %d has a duplicate ref `%s`", method, pathItemNode.Value, i, paramRefName),
					Node: param,
					Path: path,
				})
			} else {
				if paramRef != nil {
					index.paramOpRefs[pathItemNode.Value][method][paramRefName] = append(index.paramOpRefs[pathItemNode.Value][method][paramRefName], paramRef)
				}
			}

			continue

		} else {

			// param is inline.
			_, vn := utils.FindKeyNode("name", param.Content)

			path := fmt.Sprintf("$.paths['%s'].%s.parameters[%d]", pathItemNode.Value, method, i)
			if method == "top" {
				path = fmt.Sprintf("$.paths['%s'].parameters[%d]", pathItemNode.Value, i)
			}

			if vn == nil {
				index.operationParamErrors = append(index.operationParamErrors, &IndexingError{
					Err: fmt.Errorf("the '%s' operation parameter at path '%s', index %d has no 'name' value",
						method, pathItemNode.Value, i),
					Node: param,
					Path: path,
				})
				continue
			}

			ref := &Reference{
				Definition: vn.Value,
				Name:       vn.Value,
				Node:       param,
				Path:       path,
			}
			if index.paramOpRefs[pathItemNode.Value] == nil {
				index.paramOpRefs[pathItemNode.Value] = make(map[string]map[string][]*Reference)
				index.paramOpRefs[pathItemNode.Value][method] = make(map[string][]*Reference)
			}

			// if we know the path but this is a new method.
			if index.paramOpRefs[pathItemNode.Value][method] == nil {
				index.paramOpRefs[pathItemNode.Value][method] = make(map[string][]*Reference)
			}

			// if this is a duplicate name, check if the `in` type is also the same, if so, it's a duplicate.
			if len(index.paramOpRefs[pathItemNode.Value][method][ref.Name]) > 0 {

				currentNode := ref.Node
				checkNodes := index.paramOpRefs[pathItemNode.Value][method][ref.Name]
				_, currentIn := utils.FindKeyNodeTop("in", currentNode.Content)

				for _, checkNode := range checkNodes {

					_, checkIn := utils.FindKeyNodeTop("in", checkNode.Node.Content)

					if currentIn != nil && checkIn != nil && currentIn.Value == checkIn.Value {

						path := fmt.Sprintf("$.paths['%s'].%s.parameters[%d]", pathItemNode.Value, method, i)
						if method == "top" {
							path = fmt.Sprintf("$.paths['%s'].parameters[%d]", pathItemNode.Value, i)
						}

						index.operationParamErrors = append(index.operationParamErrors, &IndexingError{
							Err: fmt.Errorf("the `%s` operation parameter at path `%s`, "+
								"index %d has a duplicate name `%s` and `in` type", method, pathItemNode.Value, i, vn.Value),
							Node: param,
							Path: path,
						})
					} else {
						index.paramOpRefs[pathItemNode.Value][method][ref.Name] = append(index.paramOpRefs[pathItemNode.Value][method][ref.Name], ref)
					}
				}
			} else {
				index.paramOpRefs[pathItemNode.Value][method][ref.Name] = append(index.paramOpRefs[pathItemNode.Value][method][ref.Name], ref)
			}
			continue
		}
	}
}

func runIndexFunction(funcs []func() int, wg *sync.WaitGroup) {
	for _, cFunc := range funcs {
		go func(wg *sync.WaitGroup, cf func() int) {
			cf()
			wg.Done()
		}(wg, cFunc)
	}
}

func GenerateCleanSpecConfigBaseURL(baseURL *url.URL, dir string, includeFile bool) string {
	cleanedPath := baseURL.Path // not cleaned yet!

	// create a slice of path segments from existing path
	pathSegs := strings.Split(cleanedPath, "/")
	dirSegs := strings.Split(dir, "/")

	var cleanedSegs []string
	if !includeFile {
		dirSegs = dirSegs[:len(dirSegs)-1]
	}

	// relative paths are a pain in the ass, damn you digital ocean, use a single spec, and break them
	// down into services, please don't blast apart specs into a billion shards.
	if strings.Contains(dir, "../") {
		for s := range dirSegs {
			if dirSegs[s] == ".." {
				// chop off the last segment of the base path.
				if len(pathSegs) > 0 {
					pathSegs = pathSegs[:len(pathSegs)-1]
				}
			} else {
				cleanedSegs = append(cleanedSegs, dirSegs[s])
			}
		}
		cleanedPath = fmt.Sprintf("%s/%s", strings.Join(pathSegs, "/"), strings.Join(cleanedSegs, "/"))
	} else {
		if !strings.HasPrefix(dir, "http") {
			if len(pathSegs) > 1 || len(dirSegs) > 1 {
				cleanedPath = fmt.Sprintf("%s/%s", strings.Join(pathSegs, "/"), strings.Join(dirSegs, "/"))
			}
		} else {
			cleanedPath = strings.Join(dirSegs, "/")
		}
	}
	var p string
	if baseURL.Scheme != "" && !strings.HasPrefix(dir, "http") {
		p = fmt.Sprintf("%s://%s%s", baseURL.Scheme, baseURL.Host, cleanedPath)
	} else {
		if !strings.Contains(cleanedPath, "/") {
			p = ""
		} else {
			p = cleanedPath
		}
	}
	return strings.TrimSuffix(p, "/")
}

func syncMapToMap[K comparable, V any](sm *sync.Map) map[K]V {
	if sm == nil {
		return nil
	}

	m := make(map[K]V)

	sm.Range(func(key, value interface{}) bool {
		m[key.(K)] = value.(V)
		return true
	})

	return m
}
