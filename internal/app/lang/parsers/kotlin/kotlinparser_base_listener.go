// Code generated from KotlinParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // KotlinParser

import "github.com/antlr4-go/antlr/v4"

// BaseKotlinParserListener is a complete listener for a parse tree produced by KotlinParser.
type BaseKotlinParserListener struct{}

var _ KotlinParserListener = &BaseKotlinParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseKotlinParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseKotlinParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseKotlinParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseKotlinParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterKotlinFile is called when production kotlinFile is entered.
func (s *BaseKotlinParserListener) EnterKotlinFile(ctx *KotlinFileContext) {}

// ExitKotlinFile is called when production kotlinFile is exited.
func (s *BaseKotlinParserListener) ExitKotlinFile(ctx *KotlinFileContext) {}

// EnterScript is called when production script is entered.
func (s *BaseKotlinParserListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseKotlinParserListener) ExitScript(ctx *ScriptContext) {}

// EnterShebangLine is called when production shebangLine is entered.
func (s *BaseKotlinParserListener) EnterShebangLine(ctx *ShebangLineContext) {}

// ExitShebangLine is called when production shebangLine is exited.
func (s *BaseKotlinParserListener) ExitShebangLine(ctx *ShebangLineContext) {}

// EnterFileAnnotation is called when production fileAnnotation is entered.
func (s *BaseKotlinParserListener) EnterFileAnnotation(ctx *FileAnnotationContext) {}

// ExitFileAnnotation is called when production fileAnnotation is exited.
func (s *BaseKotlinParserListener) ExitFileAnnotation(ctx *FileAnnotationContext) {}

// EnterPackageHeader is called when production packageHeader is entered.
func (s *BaseKotlinParserListener) EnterPackageHeader(ctx *PackageHeaderContext) {}

// ExitPackageHeader is called when production packageHeader is exited.
func (s *BaseKotlinParserListener) ExitPackageHeader(ctx *PackageHeaderContext) {}

// EnterImportList is called when production importList is entered.
func (s *BaseKotlinParserListener) EnterImportList(ctx *ImportListContext) {}

// ExitImportList is called when production importList is exited.
func (s *BaseKotlinParserListener) ExitImportList(ctx *ImportListContext) {}

// EnterImportHeader is called when production importHeader is entered.
func (s *BaseKotlinParserListener) EnterImportHeader(ctx *ImportHeaderContext) {}

// ExitImportHeader is called when production importHeader is exited.
func (s *BaseKotlinParserListener) ExitImportHeader(ctx *ImportHeaderContext) {}

// EnterImportAlias is called when production importAlias is entered.
func (s *BaseKotlinParserListener) EnterImportAlias(ctx *ImportAliasContext) {}

// ExitImportAlias is called when production importAlias is exited.
func (s *BaseKotlinParserListener) ExitImportAlias(ctx *ImportAliasContext) {}

// EnterTopLevelObject is called when production topLevelObject is entered.
func (s *BaseKotlinParserListener) EnterTopLevelObject(ctx *TopLevelObjectContext) {}

// ExitTopLevelObject is called when production topLevelObject is exited.
func (s *BaseKotlinParserListener) ExitTopLevelObject(ctx *TopLevelObjectContext) {}

// EnterTypeAlias is called when production typeAlias is entered.
func (s *BaseKotlinParserListener) EnterTypeAlias(ctx *TypeAliasContext) {}

// ExitTypeAlias is called when production typeAlias is exited.
func (s *BaseKotlinParserListener) ExitTypeAlias(ctx *TypeAliasContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseKotlinParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseKotlinParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BaseKotlinParserListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BaseKotlinParserListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterPrimaryConstructor is called when production primaryConstructor is entered.
func (s *BaseKotlinParserListener) EnterPrimaryConstructor(ctx *PrimaryConstructorContext) {}

// ExitPrimaryConstructor is called when production primaryConstructor is exited.
func (s *BaseKotlinParserListener) ExitPrimaryConstructor(ctx *PrimaryConstructorContext) {}

// EnterClassBody is called when production classBody is entered.
func (s *BaseKotlinParserListener) EnterClassBody(ctx *ClassBodyContext) {}

// ExitClassBody is called when production classBody is exited.
func (s *BaseKotlinParserListener) ExitClassBody(ctx *ClassBodyContext) {}

// EnterClassParameters is called when production classParameters is entered.
func (s *BaseKotlinParserListener) EnterClassParameters(ctx *ClassParametersContext) {}

// ExitClassParameters is called when production classParameters is exited.
func (s *BaseKotlinParserListener) ExitClassParameters(ctx *ClassParametersContext) {}

// EnterClassParameter is called when production classParameter is entered.
func (s *BaseKotlinParserListener) EnterClassParameter(ctx *ClassParameterContext) {}

// ExitClassParameter is called when production classParameter is exited.
func (s *BaseKotlinParserListener) ExitClassParameter(ctx *ClassParameterContext) {}

// EnterDelegationSpecifiers is called when production delegationSpecifiers is entered.
func (s *BaseKotlinParserListener) EnterDelegationSpecifiers(ctx *DelegationSpecifiersContext) {}

// ExitDelegationSpecifiers is called when production delegationSpecifiers is exited.
func (s *BaseKotlinParserListener) ExitDelegationSpecifiers(ctx *DelegationSpecifiersContext) {}

// EnterDelegationSpecifier is called when production delegationSpecifier is entered.
func (s *BaseKotlinParserListener) EnterDelegationSpecifier(ctx *DelegationSpecifierContext) {}

// ExitDelegationSpecifier is called when production delegationSpecifier is exited.
func (s *BaseKotlinParserListener) ExitDelegationSpecifier(ctx *DelegationSpecifierContext) {}

// EnterConstructorInvocation is called when production constructorInvocation is entered.
func (s *BaseKotlinParserListener) EnterConstructorInvocation(ctx *ConstructorInvocationContext) {}

// ExitConstructorInvocation is called when production constructorInvocation is exited.
func (s *BaseKotlinParserListener) ExitConstructorInvocation(ctx *ConstructorInvocationContext) {}

// EnterAnnotatedDelegationSpecifier is called when production annotatedDelegationSpecifier is entered.
func (s *BaseKotlinParserListener) EnterAnnotatedDelegationSpecifier(ctx *AnnotatedDelegationSpecifierContext) {
}

// ExitAnnotatedDelegationSpecifier is called when production annotatedDelegationSpecifier is exited.
func (s *BaseKotlinParserListener) ExitAnnotatedDelegationSpecifier(ctx *AnnotatedDelegationSpecifierContext) {
}

// EnterExplicitDelegation is called when production explicitDelegation is entered.
func (s *BaseKotlinParserListener) EnterExplicitDelegation(ctx *ExplicitDelegationContext) {}

// ExitExplicitDelegation is called when production explicitDelegation is exited.
func (s *BaseKotlinParserListener) ExitExplicitDelegation(ctx *ExplicitDelegationContext) {}

// EnterTypeParameters is called when production typeParameters is entered.
func (s *BaseKotlinParserListener) EnterTypeParameters(ctx *TypeParametersContext) {}

// ExitTypeParameters is called when production typeParameters is exited.
func (s *BaseKotlinParserListener) ExitTypeParameters(ctx *TypeParametersContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *BaseKotlinParserListener) EnterTypeParameter(ctx *TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *BaseKotlinParserListener) ExitTypeParameter(ctx *TypeParameterContext) {}

// EnterTypeConstraints is called when production typeConstraints is entered.
func (s *BaseKotlinParserListener) EnterTypeConstraints(ctx *TypeConstraintsContext) {}

// ExitTypeConstraints is called when production typeConstraints is exited.
func (s *BaseKotlinParserListener) ExitTypeConstraints(ctx *TypeConstraintsContext) {}

// EnterTypeConstraint is called when production typeConstraint is entered.
func (s *BaseKotlinParserListener) EnterTypeConstraint(ctx *TypeConstraintContext) {}

// ExitTypeConstraint is called when production typeConstraint is exited.
func (s *BaseKotlinParserListener) ExitTypeConstraint(ctx *TypeConstraintContext) {}

// EnterClassMemberDeclarations is called when production classMemberDeclarations is entered.
func (s *BaseKotlinParserListener) EnterClassMemberDeclarations(ctx *ClassMemberDeclarationsContext) {
}

// ExitClassMemberDeclarations is called when production classMemberDeclarations is exited.
func (s *BaseKotlinParserListener) ExitClassMemberDeclarations(ctx *ClassMemberDeclarationsContext) {}

// EnterClassMemberDeclaration is called when production classMemberDeclaration is entered.
func (s *BaseKotlinParserListener) EnterClassMemberDeclaration(ctx *ClassMemberDeclarationContext) {}

// ExitClassMemberDeclaration is called when production classMemberDeclaration is exited.
func (s *BaseKotlinParserListener) ExitClassMemberDeclaration(ctx *ClassMemberDeclarationContext) {}

// EnterAnonymousInitializer is called when production anonymousInitializer is entered.
func (s *BaseKotlinParserListener) EnterAnonymousInitializer(ctx *AnonymousInitializerContext) {}

// ExitAnonymousInitializer is called when production anonymousInitializer is exited.
func (s *BaseKotlinParserListener) ExitAnonymousInitializer(ctx *AnonymousInitializerContext) {}

// EnterCompanionObject is called when production companionObject is entered.
func (s *BaseKotlinParserListener) EnterCompanionObject(ctx *CompanionObjectContext) {}

// ExitCompanionObject is called when production companionObject is exited.
func (s *BaseKotlinParserListener) ExitCompanionObject(ctx *CompanionObjectContext) {}

// EnterFunctionValueParameters is called when production functionValueParameters is entered.
func (s *BaseKotlinParserListener) EnterFunctionValueParameters(ctx *FunctionValueParametersContext) {
}

// ExitFunctionValueParameters is called when production functionValueParameters is exited.
func (s *BaseKotlinParserListener) ExitFunctionValueParameters(ctx *FunctionValueParametersContext) {}

// EnterFunctionValueParameter is called when production functionValueParameter is entered.
func (s *BaseKotlinParserListener) EnterFunctionValueParameter(ctx *FunctionValueParameterContext) {}

// ExitFunctionValueParameter is called when production functionValueParameter is exited.
func (s *BaseKotlinParserListener) ExitFunctionValueParameter(ctx *FunctionValueParameterContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseKotlinParserListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseKotlinParserListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFunctionBody is called when production functionBody is entered.
func (s *BaseKotlinParserListener) EnterFunctionBody(ctx *FunctionBodyContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *BaseKotlinParserListener) ExitFunctionBody(ctx *FunctionBodyContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseKotlinParserListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseKotlinParserListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterMultiVariableDeclaration is called when production multiVariableDeclaration is entered.
func (s *BaseKotlinParserListener) EnterMultiVariableDeclaration(ctx *MultiVariableDeclarationContext) {
}

// ExitMultiVariableDeclaration is called when production multiVariableDeclaration is exited.
func (s *BaseKotlinParserListener) ExitMultiVariableDeclaration(ctx *MultiVariableDeclarationContext) {
}

// EnterPropertyDeclaration is called when production propertyDeclaration is entered.
func (s *BaseKotlinParserListener) EnterPropertyDeclaration(ctx *PropertyDeclarationContext) {}

// ExitPropertyDeclaration is called when production propertyDeclaration is exited.
func (s *BaseKotlinParserListener) ExitPropertyDeclaration(ctx *PropertyDeclarationContext) {}

// EnterPropertyDelegate is called when production propertyDelegate is entered.
func (s *BaseKotlinParserListener) EnterPropertyDelegate(ctx *PropertyDelegateContext) {}

// ExitPropertyDelegate is called when production propertyDelegate is exited.
func (s *BaseKotlinParserListener) ExitPropertyDelegate(ctx *PropertyDelegateContext) {}

// EnterGetter is called when production getter is entered.
func (s *BaseKotlinParserListener) EnterGetter(ctx *GetterContext) {}

// ExitGetter is called when production getter is exited.
func (s *BaseKotlinParserListener) ExitGetter(ctx *GetterContext) {}

// EnterSetter is called when production setter is entered.
func (s *BaseKotlinParserListener) EnterSetter(ctx *SetterContext) {}

// ExitSetter is called when production setter is exited.
func (s *BaseKotlinParserListener) ExitSetter(ctx *SetterContext) {}

// EnterParametersWithOptionalType is called when production parametersWithOptionalType is entered.
func (s *BaseKotlinParserListener) EnterParametersWithOptionalType(ctx *ParametersWithOptionalTypeContext) {
}

// ExitParametersWithOptionalType is called when production parametersWithOptionalType is exited.
func (s *BaseKotlinParserListener) ExitParametersWithOptionalType(ctx *ParametersWithOptionalTypeContext) {
}

// EnterFunctionValueParameterWithOptionalType is called when production functionValueParameterWithOptionalType is entered.
func (s *BaseKotlinParserListener) EnterFunctionValueParameterWithOptionalType(ctx *FunctionValueParameterWithOptionalTypeContext) {
}

// ExitFunctionValueParameterWithOptionalType is called when production functionValueParameterWithOptionalType is exited.
func (s *BaseKotlinParserListener) ExitFunctionValueParameterWithOptionalType(ctx *FunctionValueParameterWithOptionalTypeContext) {
}

// EnterParameterWithOptionalType is called when production parameterWithOptionalType is entered.
func (s *BaseKotlinParserListener) EnterParameterWithOptionalType(ctx *ParameterWithOptionalTypeContext) {
}

// ExitParameterWithOptionalType is called when production parameterWithOptionalType is exited.
func (s *BaseKotlinParserListener) ExitParameterWithOptionalType(ctx *ParameterWithOptionalTypeContext) {
}

// EnterParameter is called when production parameter is entered.
func (s *BaseKotlinParserListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseKotlinParserListener) ExitParameter(ctx *ParameterContext) {}

// EnterObjectDeclaration is called when production objectDeclaration is entered.
func (s *BaseKotlinParserListener) EnterObjectDeclaration(ctx *ObjectDeclarationContext) {}

// ExitObjectDeclaration is called when production objectDeclaration is exited.
func (s *BaseKotlinParserListener) ExitObjectDeclaration(ctx *ObjectDeclarationContext) {}

// EnterSecondaryConstructor is called when production secondaryConstructor is entered.
func (s *BaseKotlinParserListener) EnterSecondaryConstructor(ctx *SecondaryConstructorContext) {}

// ExitSecondaryConstructor is called when production secondaryConstructor is exited.
func (s *BaseKotlinParserListener) ExitSecondaryConstructor(ctx *SecondaryConstructorContext) {}

// EnterConstructorDelegationCall is called when production constructorDelegationCall is entered.
func (s *BaseKotlinParserListener) EnterConstructorDelegationCall(ctx *ConstructorDelegationCallContext) {
}

// ExitConstructorDelegationCall is called when production constructorDelegationCall is exited.
func (s *BaseKotlinParserListener) ExitConstructorDelegationCall(ctx *ConstructorDelegationCallContext) {
}

// EnterEnumClassBody is called when production enumClassBody is entered.
func (s *BaseKotlinParserListener) EnterEnumClassBody(ctx *EnumClassBodyContext) {}

// ExitEnumClassBody is called when production enumClassBody is exited.
func (s *BaseKotlinParserListener) ExitEnumClassBody(ctx *EnumClassBodyContext) {}

// EnterEnumEntries is called when production enumEntries is entered.
func (s *BaseKotlinParserListener) EnterEnumEntries(ctx *EnumEntriesContext) {}

// ExitEnumEntries is called when production enumEntries is exited.
func (s *BaseKotlinParserListener) ExitEnumEntries(ctx *EnumEntriesContext) {}

// EnterEnumEntry is called when production enumEntry is entered.
func (s *BaseKotlinParserListener) EnterEnumEntry(ctx *EnumEntryContext) {}

// ExitEnumEntry is called when production enumEntry is exited.
func (s *BaseKotlinParserListener) ExitEnumEntry(ctx *EnumEntryContext) {}

// EnterType is called when production type is entered.
func (s *BaseKotlinParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseKotlinParserListener) ExitType(ctx *TypeContext) {}

// EnterTypeReference is called when production typeReference is entered.
func (s *BaseKotlinParserListener) EnterTypeReference(ctx *TypeReferenceContext) {}

// ExitTypeReference is called when production typeReference is exited.
func (s *BaseKotlinParserListener) ExitTypeReference(ctx *TypeReferenceContext) {}

// EnterNullableType is called when production nullableType is entered.
func (s *BaseKotlinParserListener) EnterNullableType(ctx *NullableTypeContext) {}

// ExitNullableType is called when production nullableType is exited.
func (s *BaseKotlinParserListener) ExitNullableType(ctx *NullableTypeContext) {}

// EnterQuest is called when production quest is entered.
func (s *BaseKotlinParserListener) EnterQuest(ctx *QuestContext) {}

// ExitQuest is called when production quest is exited.
func (s *BaseKotlinParserListener) ExitQuest(ctx *QuestContext) {}

// EnterUserType is called when production userType is entered.
func (s *BaseKotlinParserListener) EnterUserType(ctx *UserTypeContext) {}

// ExitUserType is called when production userType is exited.
func (s *BaseKotlinParserListener) ExitUserType(ctx *UserTypeContext) {}

// EnterSimpleUserType is called when production simpleUserType is entered.
func (s *BaseKotlinParserListener) EnterSimpleUserType(ctx *SimpleUserTypeContext) {}

// ExitSimpleUserType is called when production simpleUserType is exited.
func (s *BaseKotlinParserListener) ExitSimpleUserType(ctx *SimpleUserTypeContext) {}

// EnterTypeProjection is called when production typeProjection is entered.
func (s *BaseKotlinParserListener) EnterTypeProjection(ctx *TypeProjectionContext) {}

// ExitTypeProjection is called when production typeProjection is exited.
func (s *BaseKotlinParserListener) ExitTypeProjection(ctx *TypeProjectionContext) {}

// EnterTypeProjectionModifiers is called when production typeProjectionModifiers is entered.
func (s *BaseKotlinParserListener) EnterTypeProjectionModifiers(ctx *TypeProjectionModifiersContext) {
}

// ExitTypeProjectionModifiers is called when production typeProjectionModifiers is exited.
func (s *BaseKotlinParserListener) ExitTypeProjectionModifiers(ctx *TypeProjectionModifiersContext) {}

// EnterTypeProjectionModifier is called when production typeProjectionModifier is entered.
func (s *BaseKotlinParserListener) EnterTypeProjectionModifier(ctx *TypeProjectionModifierContext) {}

// ExitTypeProjectionModifier is called when production typeProjectionModifier is exited.
func (s *BaseKotlinParserListener) ExitTypeProjectionModifier(ctx *TypeProjectionModifierContext) {}

// EnterFunctionType is called when production functionType is entered.
func (s *BaseKotlinParserListener) EnterFunctionType(ctx *FunctionTypeContext) {}

// ExitFunctionType is called when production functionType is exited.
func (s *BaseKotlinParserListener) ExitFunctionType(ctx *FunctionTypeContext) {}

// EnterFunctionTypeParameters is called when production functionTypeParameters is entered.
func (s *BaseKotlinParserListener) EnterFunctionTypeParameters(ctx *FunctionTypeParametersContext) {}

// ExitFunctionTypeParameters is called when production functionTypeParameters is exited.
func (s *BaseKotlinParserListener) ExitFunctionTypeParameters(ctx *FunctionTypeParametersContext) {}

// EnterParenthesizedType is called when production parenthesizedType is entered.
func (s *BaseKotlinParserListener) EnterParenthesizedType(ctx *ParenthesizedTypeContext) {}

// ExitParenthesizedType is called when production parenthesizedType is exited.
func (s *BaseKotlinParserListener) ExitParenthesizedType(ctx *ParenthesizedTypeContext) {}

// EnterReceiverType is called when production receiverType is entered.
func (s *BaseKotlinParserListener) EnterReceiverType(ctx *ReceiverTypeContext) {}

// ExitReceiverType is called when production receiverType is exited.
func (s *BaseKotlinParserListener) ExitReceiverType(ctx *ReceiverTypeContext) {}

// EnterParenthesizedUserType is called when production parenthesizedUserType is entered.
func (s *BaseKotlinParserListener) EnterParenthesizedUserType(ctx *ParenthesizedUserTypeContext) {}

// ExitParenthesizedUserType is called when production parenthesizedUserType is exited.
func (s *BaseKotlinParserListener) ExitParenthesizedUserType(ctx *ParenthesizedUserTypeContext) {}

// EnterDefinitelyNonNullableType is called when production definitelyNonNullableType is entered.
func (s *BaseKotlinParserListener) EnterDefinitelyNonNullableType(ctx *DefinitelyNonNullableTypeContext) {
}

// ExitDefinitelyNonNullableType is called when production definitelyNonNullableType is exited.
func (s *BaseKotlinParserListener) ExitDefinitelyNonNullableType(ctx *DefinitelyNonNullableTypeContext) {
}

// EnterStatements is called when production statements is entered.
func (s *BaseKotlinParserListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaseKotlinParserListener) ExitStatements(ctx *StatementsContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseKotlinParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseKotlinParserListener) ExitStatement(ctx *StatementContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseKotlinParserListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseKotlinParserListener) ExitLabel(ctx *LabelContext) {}

// EnterControlStructureBody is called when production controlStructureBody is entered.
func (s *BaseKotlinParserListener) EnterControlStructureBody(ctx *ControlStructureBodyContext) {}

// ExitControlStructureBody is called when production controlStructureBody is exited.
func (s *BaseKotlinParserListener) ExitControlStructureBody(ctx *ControlStructureBodyContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseKotlinParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseKotlinParserListener) ExitBlock(ctx *BlockContext) {}

// EnterLoopStatement is called when production loopStatement is entered.
func (s *BaseKotlinParserListener) EnterLoopStatement(ctx *LoopStatementContext) {}

// ExitLoopStatement is called when production loopStatement is exited.
func (s *BaseKotlinParserListener) ExitLoopStatement(ctx *LoopStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseKotlinParserListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseKotlinParserListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseKotlinParserListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseKotlinParserListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterDoWhileStatement is called when production doWhileStatement is entered.
func (s *BaseKotlinParserListener) EnterDoWhileStatement(ctx *DoWhileStatementContext) {}

// ExitDoWhileStatement is called when production doWhileStatement is exited.
func (s *BaseKotlinParserListener) ExitDoWhileStatement(ctx *DoWhileStatementContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseKotlinParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseKotlinParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterSemi is called when production semi is entered.
func (s *BaseKotlinParserListener) EnterSemi(ctx *SemiContext) {}

// ExitSemi is called when production semi is exited.
func (s *BaseKotlinParserListener) ExitSemi(ctx *SemiContext) {}

// EnterSemis is called when production semis is entered.
func (s *BaseKotlinParserListener) EnterSemis(ctx *SemisContext) {}

// ExitSemis is called when production semis is exited.
func (s *BaseKotlinParserListener) ExitSemis(ctx *SemisContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseKotlinParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseKotlinParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDisjunction is called when production disjunction is entered.
func (s *BaseKotlinParserListener) EnterDisjunction(ctx *DisjunctionContext) {}

// ExitDisjunction is called when production disjunction is exited.
func (s *BaseKotlinParserListener) ExitDisjunction(ctx *DisjunctionContext) {}

// EnterConjunction is called when production conjunction is entered.
func (s *BaseKotlinParserListener) EnterConjunction(ctx *ConjunctionContext) {}

// ExitConjunction is called when production conjunction is exited.
func (s *BaseKotlinParserListener) ExitConjunction(ctx *ConjunctionContext) {}

// EnterEquality is called when production equality is entered.
func (s *BaseKotlinParserListener) EnterEquality(ctx *EqualityContext) {}

// ExitEquality is called when production equality is exited.
func (s *BaseKotlinParserListener) ExitEquality(ctx *EqualityContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseKotlinParserListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseKotlinParserListener) ExitComparison(ctx *ComparisonContext) {}

// EnterGenericCallLikeComparison is called when production genericCallLikeComparison is entered.
func (s *BaseKotlinParserListener) EnterGenericCallLikeComparison(ctx *GenericCallLikeComparisonContext) {
}

// ExitGenericCallLikeComparison is called when production genericCallLikeComparison is exited.
func (s *BaseKotlinParserListener) ExitGenericCallLikeComparison(ctx *GenericCallLikeComparisonContext) {
}

// EnterInfixOperation is called when production infixOperation is entered.
func (s *BaseKotlinParserListener) EnterInfixOperation(ctx *InfixOperationContext) {}

// ExitInfixOperation is called when production infixOperation is exited.
func (s *BaseKotlinParserListener) ExitInfixOperation(ctx *InfixOperationContext) {}

// EnterElvisExpression is called when production elvisExpression is entered.
func (s *BaseKotlinParserListener) EnterElvisExpression(ctx *ElvisExpressionContext) {}

// ExitElvisExpression is called when production elvisExpression is exited.
func (s *BaseKotlinParserListener) ExitElvisExpression(ctx *ElvisExpressionContext) {}

// EnterElvis is called when production elvis is entered.
func (s *BaseKotlinParserListener) EnterElvis(ctx *ElvisContext) {}

// ExitElvis is called when production elvis is exited.
func (s *BaseKotlinParserListener) ExitElvis(ctx *ElvisContext) {}

// EnterInfixFunctionCall is called when production infixFunctionCall is entered.
func (s *BaseKotlinParserListener) EnterInfixFunctionCall(ctx *InfixFunctionCallContext) {}

// ExitInfixFunctionCall is called when production infixFunctionCall is exited.
func (s *BaseKotlinParserListener) ExitInfixFunctionCall(ctx *InfixFunctionCallContext) {}

// EnterRangeExpression is called when production rangeExpression is entered.
func (s *BaseKotlinParserListener) EnterRangeExpression(ctx *RangeExpressionContext) {}

// ExitRangeExpression is called when production rangeExpression is exited.
func (s *BaseKotlinParserListener) ExitRangeExpression(ctx *RangeExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseKotlinParserListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseKotlinParserListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseKotlinParserListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseKotlinParserListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// EnterAsExpression is called when production asExpression is entered.
func (s *BaseKotlinParserListener) EnterAsExpression(ctx *AsExpressionContext) {}

// ExitAsExpression is called when production asExpression is exited.
func (s *BaseKotlinParserListener) ExitAsExpression(ctx *AsExpressionContext) {}

// EnterPrefixUnaryExpression is called when production prefixUnaryExpression is entered.
func (s *BaseKotlinParserListener) EnterPrefixUnaryExpression(ctx *PrefixUnaryExpressionContext) {}

// ExitPrefixUnaryExpression is called when production prefixUnaryExpression is exited.
func (s *BaseKotlinParserListener) ExitPrefixUnaryExpression(ctx *PrefixUnaryExpressionContext) {}

// EnterUnaryPrefix is called when production unaryPrefix is entered.
func (s *BaseKotlinParserListener) EnterUnaryPrefix(ctx *UnaryPrefixContext) {}

// ExitUnaryPrefix is called when production unaryPrefix is exited.
func (s *BaseKotlinParserListener) ExitUnaryPrefix(ctx *UnaryPrefixContext) {}

// EnterPostfixUnaryExpression is called when production postfixUnaryExpression is entered.
func (s *BaseKotlinParserListener) EnterPostfixUnaryExpression(ctx *PostfixUnaryExpressionContext) {}

// ExitPostfixUnaryExpression is called when production postfixUnaryExpression is exited.
func (s *BaseKotlinParserListener) ExitPostfixUnaryExpression(ctx *PostfixUnaryExpressionContext) {}

// EnterPostfixUnarySuffix is called when production postfixUnarySuffix is entered.
func (s *BaseKotlinParserListener) EnterPostfixUnarySuffix(ctx *PostfixUnarySuffixContext) {}

// ExitPostfixUnarySuffix is called when production postfixUnarySuffix is exited.
func (s *BaseKotlinParserListener) ExitPostfixUnarySuffix(ctx *PostfixUnarySuffixContext) {}

// EnterDirectlyAssignableExpression is called when production directlyAssignableExpression is entered.
func (s *BaseKotlinParserListener) EnterDirectlyAssignableExpression(ctx *DirectlyAssignableExpressionContext) {
}

// ExitDirectlyAssignableExpression is called when production directlyAssignableExpression is exited.
func (s *BaseKotlinParserListener) ExitDirectlyAssignableExpression(ctx *DirectlyAssignableExpressionContext) {
}

// EnterParenthesizedDirectlyAssignableExpression is called when production parenthesizedDirectlyAssignableExpression is entered.
func (s *BaseKotlinParserListener) EnterParenthesizedDirectlyAssignableExpression(ctx *ParenthesizedDirectlyAssignableExpressionContext) {
}

// ExitParenthesizedDirectlyAssignableExpression is called when production parenthesizedDirectlyAssignableExpression is exited.
func (s *BaseKotlinParserListener) ExitParenthesizedDirectlyAssignableExpression(ctx *ParenthesizedDirectlyAssignableExpressionContext) {
}

// EnterAssignableExpression is called when production assignableExpression is entered.
func (s *BaseKotlinParserListener) EnterAssignableExpression(ctx *AssignableExpressionContext) {}

// ExitAssignableExpression is called when production assignableExpression is exited.
func (s *BaseKotlinParserListener) ExitAssignableExpression(ctx *AssignableExpressionContext) {}

// EnterParenthesizedAssignableExpression is called when production parenthesizedAssignableExpression is entered.
func (s *BaseKotlinParserListener) EnterParenthesizedAssignableExpression(ctx *ParenthesizedAssignableExpressionContext) {
}

// ExitParenthesizedAssignableExpression is called when production parenthesizedAssignableExpression is exited.
func (s *BaseKotlinParserListener) ExitParenthesizedAssignableExpression(ctx *ParenthesizedAssignableExpressionContext) {
}

// EnterAssignableSuffix is called when production assignableSuffix is entered.
func (s *BaseKotlinParserListener) EnterAssignableSuffix(ctx *AssignableSuffixContext) {}

// ExitAssignableSuffix is called when production assignableSuffix is exited.
func (s *BaseKotlinParserListener) ExitAssignableSuffix(ctx *AssignableSuffixContext) {}

// EnterIndexingSuffix is called when production indexingSuffix is entered.
func (s *BaseKotlinParserListener) EnterIndexingSuffix(ctx *IndexingSuffixContext) {}

// ExitIndexingSuffix is called when production indexingSuffix is exited.
func (s *BaseKotlinParserListener) ExitIndexingSuffix(ctx *IndexingSuffixContext) {}

// EnterNavigationSuffix is called when production navigationSuffix is entered.
func (s *BaseKotlinParserListener) EnterNavigationSuffix(ctx *NavigationSuffixContext) {}

// ExitNavigationSuffix is called when production navigationSuffix is exited.
func (s *BaseKotlinParserListener) ExitNavigationSuffix(ctx *NavigationSuffixContext) {}

// EnterCallSuffix is called when production callSuffix is entered.
func (s *BaseKotlinParserListener) EnterCallSuffix(ctx *CallSuffixContext) {}

// ExitCallSuffix is called when production callSuffix is exited.
func (s *BaseKotlinParserListener) ExitCallSuffix(ctx *CallSuffixContext) {}

// EnterAnnotatedLambda is called when production annotatedLambda is entered.
func (s *BaseKotlinParserListener) EnterAnnotatedLambda(ctx *AnnotatedLambdaContext) {}

// ExitAnnotatedLambda is called when production annotatedLambda is exited.
func (s *BaseKotlinParserListener) ExitAnnotatedLambda(ctx *AnnotatedLambdaContext) {}

// EnterTypeArguments is called when production typeArguments is entered.
func (s *BaseKotlinParserListener) EnterTypeArguments(ctx *TypeArgumentsContext) {}

// ExitTypeArguments is called when production typeArguments is exited.
func (s *BaseKotlinParserListener) ExitTypeArguments(ctx *TypeArgumentsContext) {}

// EnterValueArguments is called when production valueArguments is entered.
func (s *BaseKotlinParserListener) EnterValueArguments(ctx *ValueArgumentsContext) {}

// ExitValueArguments is called when production valueArguments is exited.
func (s *BaseKotlinParserListener) ExitValueArguments(ctx *ValueArgumentsContext) {}

// EnterValueArgument is called when production valueArgument is entered.
func (s *BaseKotlinParserListener) EnterValueArgument(ctx *ValueArgumentContext) {}

// ExitValueArgument is called when production valueArgument is exited.
func (s *BaseKotlinParserListener) ExitValueArgument(ctx *ValueArgumentContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseKotlinParserListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseKotlinParserListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterParenthesizedExpression is called when production parenthesizedExpression is entered.
func (s *BaseKotlinParserListener) EnterParenthesizedExpression(ctx *ParenthesizedExpressionContext) {
}

// ExitParenthesizedExpression is called when production parenthesizedExpression is exited.
func (s *BaseKotlinParserListener) ExitParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// EnterCollectionLiteral is called when production collectionLiteral is entered.
func (s *BaseKotlinParserListener) EnterCollectionLiteral(ctx *CollectionLiteralContext) {}

// ExitCollectionLiteral is called when production collectionLiteral is exited.
func (s *BaseKotlinParserListener) ExitCollectionLiteral(ctx *CollectionLiteralContext) {}

// EnterLiteralConstant is called when production literalConstant is entered.
func (s *BaseKotlinParserListener) EnterLiteralConstant(ctx *LiteralConstantContext) {}

// ExitLiteralConstant is called when production literalConstant is exited.
func (s *BaseKotlinParserListener) ExitLiteralConstant(ctx *LiteralConstantContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseKotlinParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseKotlinParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterLineStringLiteral is called when production lineStringLiteral is entered.
func (s *BaseKotlinParserListener) EnterLineStringLiteral(ctx *LineStringLiteralContext) {}

// ExitLineStringLiteral is called when production lineStringLiteral is exited.
func (s *BaseKotlinParserListener) ExitLineStringLiteral(ctx *LineStringLiteralContext) {}

// EnterMultiLineStringLiteral is called when production multiLineStringLiteral is entered.
func (s *BaseKotlinParserListener) EnterMultiLineStringLiteral(ctx *MultiLineStringLiteralContext) {}

// ExitMultiLineStringLiteral is called when production multiLineStringLiteral is exited.
func (s *BaseKotlinParserListener) ExitMultiLineStringLiteral(ctx *MultiLineStringLiteralContext) {}

// EnterLineStringContent is called when production lineStringContent is entered.
func (s *BaseKotlinParserListener) EnterLineStringContent(ctx *LineStringContentContext) {}

// ExitLineStringContent is called when production lineStringContent is exited.
func (s *BaseKotlinParserListener) ExitLineStringContent(ctx *LineStringContentContext) {}

// EnterLineStringExpression is called when production lineStringExpression is entered.
func (s *BaseKotlinParserListener) EnterLineStringExpression(ctx *LineStringExpressionContext) {}

// ExitLineStringExpression is called when production lineStringExpression is exited.
func (s *BaseKotlinParserListener) ExitLineStringExpression(ctx *LineStringExpressionContext) {}

// EnterMultiLineStringContent is called when production multiLineStringContent is entered.
func (s *BaseKotlinParserListener) EnterMultiLineStringContent(ctx *MultiLineStringContentContext) {}

// ExitMultiLineStringContent is called when production multiLineStringContent is exited.
func (s *BaseKotlinParserListener) ExitMultiLineStringContent(ctx *MultiLineStringContentContext) {}

// EnterMultiLineStringExpression is called when production multiLineStringExpression is entered.
func (s *BaseKotlinParserListener) EnterMultiLineStringExpression(ctx *MultiLineStringExpressionContext) {
}

// ExitMultiLineStringExpression is called when production multiLineStringExpression is exited.
func (s *BaseKotlinParserListener) ExitMultiLineStringExpression(ctx *MultiLineStringExpressionContext) {
}

// EnterLambdaLiteral is called when production lambdaLiteral is entered.
func (s *BaseKotlinParserListener) EnterLambdaLiteral(ctx *LambdaLiteralContext) {}

// ExitLambdaLiteral is called when production lambdaLiteral is exited.
func (s *BaseKotlinParserListener) ExitLambdaLiteral(ctx *LambdaLiteralContext) {}

// EnterLambdaParameters is called when production lambdaParameters is entered.
func (s *BaseKotlinParserListener) EnterLambdaParameters(ctx *LambdaParametersContext) {}

// ExitLambdaParameters is called when production lambdaParameters is exited.
func (s *BaseKotlinParserListener) ExitLambdaParameters(ctx *LambdaParametersContext) {}

// EnterLambdaParameter is called when production lambdaParameter is entered.
func (s *BaseKotlinParserListener) EnterLambdaParameter(ctx *LambdaParameterContext) {}

// ExitLambdaParameter is called when production lambdaParameter is exited.
func (s *BaseKotlinParserListener) ExitLambdaParameter(ctx *LambdaParameterContext) {}

// EnterAnonymousFunction is called when production anonymousFunction is entered.
func (s *BaseKotlinParserListener) EnterAnonymousFunction(ctx *AnonymousFunctionContext) {}

// ExitAnonymousFunction is called when production anonymousFunction is exited.
func (s *BaseKotlinParserListener) ExitAnonymousFunction(ctx *AnonymousFunctionContext) {}

// EnterFunctionLiteral is called when production functionLiteral is entered.
func (s *BaseKotlinParserListener) EnterFunctionLiteral(ctx *FunctionLiteralContext) {}

// ExitFunctionLiteral is called when production functionLiteral is exited.
func (s *BaseKotlinParserListener) ExitFunctionLiteral(ctx *FunctionLiteralContext) {}

// EnterObjectLiteral is called when production objectLiteral is entered.
func (s *BaseKotlinParserListener) EnterObjectLiteral(ctx *ObjectLiteralContext) {}

// ExitObjectLiteral is called when production objectLiteral is exited.
func (s *BaseKotlinParserListener) ExitObjectLiteral(ctx *ObjectLiteralContext) {}

// EnterThisExpression is called when production thisExpression is entered.
func (s *BaseKotlinParserListener) EnterThisExpression(ctx *ThisExpressionContext) {}

// ExitThisExpression is called when production thisExpression is exited.
func (s *BaseKotlinParserListener) ExitThisExpression(ctx *ThisExpressionContext) {}

// EnterSuperExpression is called when production superExpression is entered.
func (s *BaseKotlinParserListener) EnterSuperExpression(ctx *SuperExpressionContext) {}

// ExitSuperExpression is called when production superExpression is exited.
func (s *BaseKotlinParserListener) ExitSuperExpression(ctx *SuperExpressionContext) {}

// EnterIfExpression is called when production ifExpression is entered.
func (s *BaseKotlinParserListener) EnterIfExpression(ctx *IfExpressionContext) {}

// ExitIfExpression is called when production ifExpression is exited.
func (s *BaseKotlinParserListener) ExitIfExpression(ctx *IfExpressionContext) {}

// EnterWhenSubject is called when production whenSubject is entered.
func (s *BaseKotlinParserListener) EnterWhenSubject(ctx *WhenSubjectContext) {}

// ExitWhenSubject is called when production whenSubject is exited.
func (s *BaseKotlinParserListener) ExitWhenSubject(ctx *WhenSubjectContext) {}

// EnterWhenExpression is called when production whenExpression is entered.
func (s *BaseKotlinParserListener) EnterWhenExpression(ctx *WhenExpressionContext) {}

// ExitWhenExpression is called when production whenExpression is exited.
func (s *BaseKotlinParserListener) ExitWhenExpression(ctx *WhenExpressionContext) {}

// EnterWhenEntry is called when production whenEntry is entered.
func (s *BaseKotlinParserListener) EnterWhenEntry(ctx *WhenEntryContext) {}

// ExitWhenEntry is called when production whenEntry is exited.
func (s *BaseKotlinParserListener) ExitWhenEntry(ctx *WhenEntryContext) {}

// EnterWhenCondition is called when production whenCondition is entered.
func (s *BaseKotlinParserListener) EnterWhenCondition(ctx *WhenConditionContext) {}

// ExitWhenCondition is called when production whenCondition is exited.
func (s *BaseKotlinParserListener) ExitWhenCondition(ctx *WhenConditionContext) {}

// EnterRangeTest is called when production rangeTest is entered.
func (s *BaseKotlinParserListener) EnterRangeTest(ctx *RangeTestContext) {}

// ExitRangeTest is called when production rangeTest is exited.
func (s *BaseKotlinParserListener) ExitRangeTest(ctx *RangeTestContext) {}

// EnterTypeTest is called when production typeTest is entered.
func (s *BaseKotlinParserListener) EnterTypeTest(ctx *TypeTestContext) {}

// ExitTypeTest is called when production typeTest is exited.
func (s *BaseKotlinParserListener) ExitTypeTest(ctx *TypeTestContext) {}

// EnterTryExpression is called when production tryExpression is entered.
func (s *BaseKotlinParserListener) EnterTryExpression(ctx *TryExpressionContext) {}

// ExitTryExpression is called when production tryExpression is exited.
func (s *BaseKotlinParserListener) ExitTryExpression(ctx *TryExpressionContext) {}

// EnterCatchBlock is called when production catchBlock is entered.
func (s *BaseKotlinParserListener) EnterCatchBlock(ctx *CatchBlockContext) {}

// ExitCatchBlock is called when production catchBlock is exited.
func (s *BaseKotlinParserListener) ExitCatchBlock(ctx *CatchBlockContext) {}

// EnterFinallyBlock is called when production finallyBlock is entered.
func (s *BaseKotlinParserListener) EnterFinallyBlock(ctx *FinallyBlockContext) {}

// ExitFinallyBlock is called when production finallyBlock is exited.
func (s *BaseKotlinParserListener) ExitFinallyBlock(ctx *FinallyBlockContext) {}

// EnterJumpExpression is called when production jumpExpression is entered.
func (s *BaseKotlinParserListener) EnterJumpExpression(ctx *JumpExpressionContext) {}

// ExitJumpExpression is called when production jumpExpression is exited.
func (s *BaseKotlinParserListener) ExitJumpExpression(ctx *JumpExpressionContext) {}

// EnterCallableReference is called when production callableReference is entered.
func (s *BaseKotlinParserListener) EnterCallableReference(ctx *CallableReferenceContext) {}

// ExitCallableReference is called when production callableReference is exited.
func (s *BaseKotlinParserListener) ExitCallableReference(ctx *CallableReferenceContext) {}

// EnterAssignmentAndOperator is called when production assignmentAndOperator is entered.
func (s *BaseKotlinParserListener) EnterAssignmentAndOperator(ctx *AssignmentAndOperatorContext) {}

// ExitAssignmentAndOperator is called when production assignmentAndOperator is exited.
func (s *BaseKotlinParserListener) ExitAssignmentAndOperator(ctx *AssignmentAndOperatorContext) {}

// EnterEqualityOperator is called when production equalityOperator is entered.
func (s *BaseKotlinParserListener) EnterEqualityOperator(ctx *EqualityOperatorContext) {}

// ExitEqualityOperator is called when production equalityOperator is exited.
func (s *BaseKotlinParserListener) ExitEqualityOperator(ctx *EqualityOperatorContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseKotlinParserListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseKotlinParserListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterInOperator is called when production inOperator is entered.
func (s *BaseKotlinParserListener) EnterInOperator(ctx *InOperatorContext) {}

// ExitInOperator is called when production inOperator is exited.
func (s *BaseKotlinParserListener) ExitInOperator(ctx *InOperatorContext) {}

// EnterIsOperator is called when production isOperator is entered.
func (s *BaseKotlinParserListener) EnterIsOperator(ctx *IsOperatorContext) {}

// ExitIsOperator is called when production isOperator is exited.
func (s *BaseKotlinParserListener) ExitIsOperator(ctx *IsOperatorContext) {}

// EnterAdditiveOperator is called when production additiveOperator is entered.
func (s *BaseKotlinParserListener) EnterAdditiveOperator(ctx *AdditiveOperatorContext) {}

// ExitAdditiveOperator is called when production additiveOperator is exited.
func (s *BaseKotlinParserListener) ExitAdditiveOperator(ctx *AdditiveOperatorContext) {}

// EnterMultiplicativeOperator is called when production multiplicativeOperator is entered.
func (s *BaseKotlinParserListener) EnterMultiplicativeOperator(ctx *MultiplicativeOperatorContext) {}

// ExitMultiplicativeOperator is called when production multiplicativeOperator is exited.
func (s *BaseKotlinParserListener) ExitMultiplicativeOperator(ctx *MultiplicativeOperatorContext) {}

// EnterAsOperator is called when production asOperator is entered.
func (s *BaseKotlinParserListener) EnterAsOperator(ctx *AsOperatorContext) {}

// ExitAsOperator is called when production asOperator is exited.
func (s *BaseKotlinParserListener) ExitAsOperator(ctx *AsOperatorContext) {}

// EnterPrefixUnaryOperator is called when production prefixUnaryOperator is entered.
func (s *BaseKotlinParserListener) EnterPrefixUnaryOperator(ctx *PrefixUnaryOperatorContext) {}

// ExitPrefixUnaryOperator is called when production prefixUnaryOperator is exited.
func (s *BaseKotlinParserListener) ExitPrefixUnaryOperator(ctx *PrefixUnaryOperatorContext) {}

// EnterPostfixUnaryOperator is called when production postfixUnaryOperator is entered.
func (s *BaseKotlinParserListener) EnterPostfixUnaryOperator(ctx *PostfixUnaryOperatorContext) {}

// ExitPostfixUnaryOperator is called when production postfixUnaryOperator is exited.
func (s *BaseKotlinParserListener) ExitPostfixUnaryOperator(ctx *PostfixUnaryOperatorContext) {}

// EnterExcl is called when production excl is entered.
func (s *BaseKotlinParserListener) EnterExcl(ctx *ExclContext) {}

// ExitExcl is called when production excl is exited.
func (s *BaseKotlinParserListener) ExitExcl(ctx *ExclContext) {}

// EnterMemberAccessOperator is called when production memberAccessOperator is entered.
func (s *BaseKotlinParserListener) EnterMemberAccessOperator(ctx *MemberAccessOperatorContext) {}

// ExitMemberAccessOperator is called when production memberAccessOperator is exited.
func (s *BaseKotlinParserListener) ExitMemberAccessOperator(ctx *MemberAccessOperatorContext) {}

// EnterSafeNav is called when production safeNav is entered.
func (s *BaseKotlinParserListener) EnterSafeNav(ctx *SafeNavContext) {}

// ExitSafeNav is called when production safeNav is exited.
func (s *BaseKotlinParserListener) ExitSafeNav(ctx *SafeNavContext) {}

// EnterModifiers is called when production modifiers is entered.
func (s *BaseKotlinParserListener) EnterModifiers(ctx *ModifiersContext) {}

// ExitModifiers is called when production modifiers is exited.
func (s *BaseKotlinParserListener) ExitModifiers(ctx *ModifiersContext) {}

// EnterParameterModifiers is called when production parameterModifiers is entered.
func (s *BaseKotlinParserListener) EnterParameterModifiers(ctx *ParameterModifiersContext) {}

// ExitParameterModifiers is called when production parameterModifiers is exited.
func (s *BaseKotlinParserListener) ExitParameterModifiers(ctx *ParameterModifiersContext) {}

// EnterModifier is called when production modifier is entered.
func (s *BaseKotlinParserListener) EnterModifier(ctx *ModifierContext) {}

// ExitModifier is called when production modifier is exited.
func (s *BaseKotlinParserListener) ExitModifier(ctx *ModifierContext) {}

// EnterTypeModifiers is called when production typeModifiers is entered.
func (s *BaseKotlinParserListener) EnterTypeModifiers(ctx *TypeModifiersContext) {}

// ExitTypeModifiers is called when production typeModifiers is exited.
func (s *BaseKotlinParserListener) ExitTypeModifiers(ctx *TypeModifiersContext) {}

// EnterTypeModifier is called when production typeModifier is entered.
func (s *BaseKotlinParserListener) EnterTypeModifier(ctx *TypeModifierContext) {}

// ExitTypeModifier is called when production typeModifier is exited.
func (s *BaseKotlinParserListener) ExitTypeModifier(ctx *TypeModifierContext) {}

// EnterClassModifier is called when production classModifier is entered.
func (s *BaseKotlinParserListener) EnterClassModifier(ctx *ClassModifierContext) {}

// ExitClassModifier is called when production classModifier is exited.
func (s *BaseKotlinParserListener) ExitClassModifier(ctx *ClassModifierContext) {}

// EnterMemberModifier is called when production memberModifier is entered.
func (s *BaseKotlinParserListener) EnterMemberModifier(ctx *MemberModifierContext) {}

// ExitMemberModifier is called when production memberModifier is exited.
func (s *BaseKotlinParserListener) ExitMemberModifier(ctx *MemberModifierContext) {}

// EnterVisibilityModifier is called when production visibilityModifier is entered.
func (s *BaseKotlinParserListener) EnterVisibilityModifier(ctx *VisibilityModifierContext) {}

// ExitVisibilityModifier is called when production visibilityModifier is exited.
func (s *BaseKotlinParserListener) ExitVisibilityModifier(ctx *VisibilityModifierContext) {}

// EnterVarianceModifier is called when production varianceModifier is entered.
func (s *BaseKotlinParserListener) EnterVarianceModifier(ctx *VarianceModifierContext) {}

// ExitVarianceModifier is called when production varianceModifier is exited.
func (s *BaseKotlinParserListener) ExitVarianceModifier(ctx *VarianceModifierContext) {}

// EnterTypeParameterModifiers is called when production typeParameterModifiers is entered.
func (s *BaseKotlinParserListener) EnterTypeParameterModifiers(ctx *TypeParameterModifiersContext) {}

// ExitTypeParameterModifiers is called when production typeParameterModifiers is exited.
func (s *BaseKotlinParserListener) ExitTypeParameterModifiers(ctx *TypeParameterModifiersContext) {}

// EnterTypeParameterModifier is called when production typeParameterModifier is entered.
func (s *BaseKotlinParserListener) EnterTypeParameterModifier(ctx *TypeParameterModifierContext) {}

// ExitTypeParameterModifier is called when production typeParameterModifier is exited.
func (s *BaseKotlinParserListener) ExitTypeParameterModifier(ctx *TypeParameterModifierContext) {}

// EnterFunctionModifier is called when production functionModifier is entered.
func (s *BaseKotlinParserListener) EnterFunctionModifier(ctx *FunctionModifierContext) {}

// ExitFunctionModifier is called when production functionModifier is exited.
func (s *BaseKotlinParserListener) ExitFunctionModifier(ctx *FunctionModifierContext) {}

// EnterPropertyModifier is called when production propertyModifier is entered.
func (s *BaseKotlinParserListener) EnterPropertyModifier(ctx *PropertyModifierContext) {}

// ExitPropertyModifier is called when production propertyModifier is exited.
func (s *BaseKotlinParserListener) ExitPropertyModifier(ctx *PropertyModifierContext) {}

// EnterInheritanceModifier is called when production inheritanceModifier is entered.
func (s *BaseKotlinParserListener) EnterInheritanceModifier(ctx *InheritanceModifierContext) {}

// ExitInheritanceModifier is called when production inheritanceModifier is exited.
func (s *BaseKotlinParserListener) ExitInheritanceModifier(ctx *InheritanceModifierContext) {}

// EnterParameterModifier is called when production parameterModifier is entered.
func (s *BaseKotlinParserListener) EnterParameterModifier(ctx *ParameterModifierContext) {}

// ExitParameterModifier is called when production parameterModifier is exited.
func (s *BaseKotlinParserListener) ExitParameterModifier(ctx *ParameterModifierContext) {}

// EnterReificationModifier is called when production reificationModifier is entered.
func (s *BaseKotlinParserListener) EnterReificationModifier(ctx *ReificationModifierContext) {}

// ExitReificationModifier is called when production reificationModifier is exited.
func (s *BaseKotlinParserListener) ExitReificationModifier(ctx *ReificationModifierContext) {}

// EnterPlatformModifier is called when production platformModifier is entered.
func (s *BaseKotlinParserListener) EnterPlatformModifier(ctx *PlatformModifierContext) {}

// ExitPlatformModifier is called when production platformModifier is exited.
func (s *BaseKotlinParserListener) ExitPlatformModifier(ctx *PlatformModifierContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseKotlinParserListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseKotlinParserListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterSingleAnnotation is called when production singleAnnotation is entered.
func (s *BaseKotlinParserListener) EnterSingleAnnotation(ctx *SingleAnnotationContext) {}

// ExitSingleAnnotation is called when production singleAnnotation is exited.
func (s *BaseKotlinParserListener) ExitSingleAnnotation(ctx *SingleAnnotationContext) {}

// EnterMultiAnnotation is called when production multiAnnotation is entered.
func (s *BaseKotlinParserListener) EnterMultiAnnotation(ctx *MultiAnnotationContext) {}

// ExitMultiAnnotation is called when production multiAnnotation is exited.
func (s *BaseKotlinParserListener) ExitMultiAnnotation(ctx *MultiAnnotationContext) {}

// EnterAnnotationUseSiteTarget is called when production annotationUseSiteTarget is entered.
func (s *BaseKotlinParserListener) EnterAnnotationUseSiteTarget(ctx *AnnotationUseSiteTargetContext) {
}

// ExitAnnotationUseSiteTarget is called when production annotationUseSiteTarget is exited.
func (s *BaseKotlinParserListener) ExitAnnotationUseSiteTarget(ctx *AnnotationUseSiteTargetContext) {}

// EnterUnescapedAnnotation is called when production unescapedAnnotation is entered.
func (s *BaseKotlinParserListener) EnterUnescapedAnnotation(ctx *UnescapedAnnotationContext) {}

// ExitUnescapedAnnotation is called when production unescapedAnnotation is exited.
func (s *BaseKotlinParserListener) ExitUnescapedAnnotation(ctx *UnescapedAnnotationContext) {}

// EnterSimpleIdentifier is called when production simpleIdentifier is entered.
func (s *BaseKotlinParserListener) EnterSimpleIdentifier(ctx *SimpleIdentifierContext) {}

// ExitSimpleIdentifier is called when production simpleIdentifier is exited.
func (s *BaseKotlinParserListener) ExitSimpleIdentifier(ctx *SimpleIdentifierContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseKotlinParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseKotlinParserListener) ExitIdentifier(ctx *IdentifierContext) {}
