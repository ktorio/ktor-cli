// Code generated from grammars/kotlin/KotlinParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // KotlinParser

import "github.com/antlr4-go/antlr/v4"

// KotlinParserListener is a complete listener for a parse tree produced by KotlinParser.
type KotlinParserListener interface {
	antlr.ParseTreeListener

	// EnterKotlinFile is called when entering the kotlinFile production.
	EnterKotlinFile(c *KotlinFileContext)

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// EnterShebangLine is called when entering the shebangLine production.
	EnterShebangLine(c *ShebangLineContext)

	// EnterFileAnnotation is called when entering the fileAnnotation production.
	EnterFileAnnotation(c *FileAnnotationContext)

	// EnterPackageHeader is called when entering the packageHeader production.
	EnterPackageHeader(c *PackageHeaderContext)

	// EnterImportList is called when entering the importList production.
	EnterImportList(c *ImportListContext)

	// EnterImportHeader is called when entering the importHeader production.
	EnterImportHeader(c *ImportHeaderContext)

	// EnterImportAlias is called when entering the importAlias production.
	EnterImportAlias(c *ImportAliasContext)

	// EnterTopLevelObject is called when entering the topLevelObject production.
	EnterTopLevelObject(c *TopLevelObjectContext)

	// EnterTypeAlias is called when entering the typeAlias production.
	EnterTypeAlias(c *TypeAliasContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterPrimaryConstructor is called when entering the primaryConstructor production.
	EnterPrimaryConstructor(c *PrimaryConstructorContext)

	// EnterClassBody is called when entering the classBody production.
	EnterClassBody(c *ClassBodyContext)

	// EnterClassParameters is called when entering the classParameters production.
	EnterClassParameters(c *ClassParametersContext)

	// EnterClassParameter is called when entering the classParameter production.
	EnterClassParameter(c *ClassParameterContext)

	// EnterDelegationSpecifiers is called when entering the delegationSpecifiers production.
	EnterDelegationSpecifiers(c *DelegationSpecifiersContext)

	// EnterDelegationSpecifier is called when entering the delegationSpecifier production.
	EnterDelegationSpecifier(c *DelegationSpecifierContext)

	// EnterConstructorInvocation is called when entering the constructorInvocation production.
	EnterConstructorInvocation(c *ConstructorInvocationContext)

	// EnterAnnotatedDelegationSpecifier is called when entering the annotatedDelegationSpecifier production.
	EnterAnnotatedDelegationSpecifier(c *AnnotatedDelegationSpecifierContext)

	// EnterExplicitDelegation is called when entering the explicitDelegation production.
	EnterExplicitDelegation(c *ExplicitDelegationContext)

	// EnterTypeParameters is called when entering the typeParameters production.
	EnterTypeParameters(c *TypeParametersContext)

	// EnterTypeParameter is called when entering the typeParameter production.
	EnterTypeParameter(c *TypeParameterContext)

	// EnterTypeConstraints is called when entering the typeConstraints production.
	EnterTypeConstraints(c *TypeConstraintsContext)

	// EnterTypeConstraint is called when entering the typeConstraint production.
	EnterTypeConstraint(c *TypeConstraintContext)

	// EnterClassMemberDeclarations is called when entering the classMemberDeclarations production.
	EnterClassMemberDeclarations(c *ClassMemberDeclarationsContext)

	// EnterClassMemberDeclaration is called when entering the classMemberDeclaration production.
	EnterClassMemberDeclaration(c *ClassMemberDeclarationContext)

	// EnterAnonymousInitializer is called when entering the anonymousInitializer production.
	EnterAnonymousInitializer(c *AnonymousInitializerContext)

	// EnterCompanionObject is called when entering the companionObject production.
	EnterCompanionObject(c *CompanionObjectContext)

	// EnterFunctionValueParameters is called when entering the functionValueParameters production.
	EnterFunctionValueParameters(c *FunctionValueParametersContext)

	// EnterFunctionValueParameter is called when entering the functionValueParameter production.
	EnterFunctionValueParameter(c *FunctionValueParameterContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFunctionBody is called when entering the functionBody production.
	EnterFunctionBody(c *FunctionBodyContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterMultiVariableDeclaration is called when entering the multiVariableDeclaration production.
	EnterMultiVariableDeclaration(c *MultiVariableDeclarationContext)

	// EnterPropertyDeclaration is called when entering the propertyDeclaration production.
	EnterPropertyDeclaration(c *PropertyDeclarationContext)

	// EnterPropertyDelegate is called when entering the propertyDelegate production.
	EnterPropertyDelegate(c *PropertyDelegateContext)

	// EnterGetter is called when entering the getter production.
	EnterGetter(c *GetterContext)

	// EnterSetter is called when entering the setter production.
	EnterSetter(c *SetterContext)

	// EnterParametersWithOptionalType is called when entering the parametersWithOptionalType production.
	EnterParametersWithOptionalType(c *ParametersWithOptionalTypeContext)

	// EnterFunctionValueParameterWithOptionalType is called when entering the functionValueParameterWithOptionalType production.
	EnterFunctionValueParameterWithOptionalType(c *FunctionValueParameterWithOptionalTypeContext)

	// EnterParameterWithOptionalType is called when entering the parameterWithOptionalType production.
	EnterParameterWithOptionalType(c *ParameterWithOptionalTypeContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterObjectDeclaration is called when entering the objectDeclaration production.
	EnterObjectDeclaration(c *ObjectDeclarationContext)

	// EnterSecondaryConstructor is called when entering the secondaryConstructor production.
	EnterSecondaryConstructor(c *SecondaryConstructorContext)

	// EnterConstructorDelegationCall is called when entering the constructorDelegationCall production.
	EnterConstructorDelegationCall(c *ConstructorDelegationCallContext)

	// EnterEnumClassBody is called when entering the enumClassBody production.
	EnterEnumClassBody(c *EnumClassBodyContext)

	// EnterEnumEntries is called when entering the enumEntries production.
	EnterEnumEntries(c *EnumEntriesContext)

	// EnterEnumEntry is called when entering the enumEntry production.
	EnterEnumEntry(c *EnumEntryContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterTypeReference is called when entering the typeReference production.
	EnterTypeReference(c *TypeReferenceContext)

	// EnterNullableType is called when entering the nullableType production.
	EnterNullableType(c *NullableTypeContext)

	// EnterQuest is called when entering the quest production.
	EnterQuest(c *QuestContext)

	// EnterUserType is called when entering the userType production.
	EnterUserType(c *UserTypeContext)

	// EnterSimpleUserType is called when entering the simpleUserType production.
	EnterSimpleUserType(c *SimpleUserTypeContext)

	// EnterTypeProjection is called when entering the typeProjection production.
	EnterTypeProjection(c *TypeProjectionContext)

	// EnterTypeProjectionModifiers is called when entering the typeProjectionModifiers production.
	EnterTypeProjectionModifiers(c *TypeProjectionModifiersContext)

	// EnterTypeProjectionModifier is called when entering the typeProjectionModifier production.
	EnterTypeProjectionModifier(c *TypeProjectionModifierContext)

	// EnterFunctionType is called when entering the functionType production.
	EnterFunctionType(c *FunctionTypeContext)

	// EnterFunctionTypeParameters is called when entering the functionTypeParameters production.
	EnterFunctionTypeParameters(c *FunctionTypeParametersContext)

	// EnterParenthesizedType is called when entering the parenthesizedType production.
	EnterParenthesizedType(c *ParenthesizedTypeContext)

	// EnterReceiverType is called when entering the receiverType production.
	EnterReceiverType(c *ReceiverTypeContext)

	// EnterParenthesizedUserType is called when entering the parenthesizedUserType production.
	EnterParenthesizedUserType(c *ParenthesizedUserTypeContext)

	// EnterDefinitelyNonNullableType is called when entering the definitelyNonNullableType production.
	EnterDefinitelyNonNullableType(c *DefinitelyNonNullableTypeContext)

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterControlStructureBody is called when entering the controlStructureBody production.
	EnterControlStructureBody(c *ControlStructureBodyContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterLoopStatement is called when entering the loopStatement production.
	EnterLoopStatement(c *LoopStatementContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterDoWhileStatement is called when entering the doWhileStatement production.
	EnterDoWhileStatement(c *DoWhileStatementContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterSemi is called when entering the semi production.
	EnterSemi(c *SemiContext)

	// EnterSemis is called when entering the semis production.
	EnterSemis(c *SemisContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterInfixFunctionCall is called when entering the infixFunctionCall production.
	EnterInfixFunctionCall(c *InfixFunctionCallContext)

	// EnterPrefixUnaryExpression is called when entering the prefixUnaryExpression production.
	EnterPrefixUnaryExpression(c *PrefixUnaryExpressionContext)

	// EnterUnaryPrefix is called when entering the unaryPrefix production.
	EnterUnaryPrefix(c *UnaryPrefixContext)

	// EnterPostfixUnaryExpression is called when entering the postfixUnaryExpression production.
	EnterPostfixUnaryExpression(c *PostfixUnaryExpressionContext)

	// EnterPostfixUnarySuffix is called when entering the postfixUnarySuffix production.
	EnterPostfixUnarySuffix(c *PostfixUnarySuffixContext)

	// EnterDirectlyAssignableExpression is called when entering the directlyAssignableExpression production.
	EnterDirectlyAssignableExpression(c *DirectlyAssignableExpressionContext)

	// EnterParenthesizedDirectlyAssignableExpression is called when entering the parenthesizedDirectlyAssignableExpression production.
	EnterParenthesizedDirectlyAssignableExpression(c *ParenthesizedDirectlyAssignableExpressionContext)

	// EnterAssignableExpression is called when entering the assignableExpression production.
	EnterAssignableExpression(c *AssignableExpressionContext)

	// EnterParenthesizedAssignableExpression is called when entering the parenthesizedAssignableExpression production.
	EnterParenthesizedAssignableExpression(c *ParenthesizedAssignableExpressionContext)

	// EnterAssignableSuffix is called when entering the assignableSuffix production.
	EnterAssignableSuffix(c *AssignableSuffixContext)

	// EnterIndexingSuffix is called when entering the indexingSuffix production.
	EnterIndexingSuffix(c *IndexingSuffixContext)

	// EnterNavigationSuffix is called when entering the navigationSuffix production.
	EnterNavigationSuffix(c *NavigationSuffixContext)

	// EnterCallSuffix is called when entering the callSuffix production.
	EnterCallSuffix(c *CallSuffixContext)

	// EnterAnnotatedLambda is called when entering the annotatedLambda production.
	EnterAnnotatedLambda(c *AnnotatedLambdaContext)

	// EnterTypeArguments is called when entering the typeArguments production.
	EnterTypeArguments(c *TypeArgumentsContext)

	// EnterValueArguments is called when entering the valueArguments production.
	EnterValueArguments(c *ValueArgumentsContext)

	// EnterValueArgument is called when entering the valueArgument production.
	EnterValueArgument(c *ValueArgumentContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterParenthesizedExpression is called when entering the parenthesizedExpression production.
	EnterParenthesizedExpression(c *ParenthesizedExpressionContext)

	// EnterCollectionLiteral is called when entering the collectionLiteral production.
	EnterCollectionLiteral(c *CollectionLiteralContext)

	// EnterLiteralConstant is called when entering the literalConstant production.
	EnterLiteralConstant(c *LiteralConstantContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterLineStringLiteral is called when entering the lineStringLiteral production.
	EnterLineStringLiteral(c *LineStringLiteralContext)

	// EnterMultiLineStringLiteral is called when entering the multiLineStringLiteral production.
	EnterMultiLineStringLiteral(c *MultiLineStringLiteralContext)

	// EnterLineStringContent is called when entering the lineStringContent production.
	EnterLineStringContent(c *LineStringContentContext)

	// EnterLineStringExpression is called when entering the lineStringExpression production.
	EnterLineStringExpression(c *LineStringExpressionContext)

	// EnterMultiLineStringContent is called when entering the multiLineStringContent production.
	EnterMultiLineStringContent(c *MultiLineStringContentContext)

	// EnterMultiLineStringExpression is called when entering the multiLineStringExpression production.
	EnterMultiLineStringExpression(c *MultiLineStringExpressionContext)

	// EnterLambdaLiteral is called when entering the lambdaLiteral production.
	EnterLambdaLiteral(c *LambdaLiteralContext)

	// EnterLambdaParameters is called when entering the lambdaParameters production.
	EnterLambdaParameters(c *LambdaParametersContext)

	// EnterLambdaParameter is called when entering the lambdaParameter production.
	EnterLambdaParameter(c *LambdaParameterContext)

	// EnterAnonymousFunction is called when entering the anonymousFunction production.
	EnterAnonymousFunction(c *AnonymousFunctionContext)

	// EnterFunctionLiteral is called when entering the functionLiteral production.
	EnterFunctionLiteral(c *FunctionLiteralContext)

	// EnterObjectLiteral is called when entering the objectLiteral production.
	EnterObjectLiteral(c *ObjectLiteralContext)

	// EnterThisExpression is called when entering the thisExpression production.
	EnterThisExpression(c *ThisExpressionContext)

	// EnterSuperExpression is called when entering the superExpression production.
	EnterSuperExpression(c *SuperExpressionContext)

	// EnterIfExpression is called when entering the ifExpression production.
	EnterIfExpression(c *IfExpressionContext)

	// EnterWhenSubject is called when entering the whenSubject production.
	EnterWhenSubject(c *WhenSubjectContext)

	// EnterWhenExpression is called when entering the whenExpression production.
	EnterWhenExpression(c *WhenExpressionContext)

	// EnterWhenEntry is called when entering the whenEntry production.
	EnterWhenEntry(c *WhenEntryContext)

	// EnterWhenCondition is called when entering the whenCondition production.
	EnterWhenCondition(c *WhenConditionContext)

	// EnterRangeTest is called when entering the rangeTest production.
	EnterRangeTest(c *RangeTestContext)

	// EnterTypeTest is called when entering the typeTest production.
	EnterTypeTest(c *TypeTestContext)

	// EnterTryExpression is called when entering the tryExpression production.
	EnterTryExpression(c *TryExpressionContext)

	// EnterCatchBlock is called when entering the catchBlock production.
	EnterCatchBlock(c *CatchBlockContext)

	// EnterFinallyBlock is called when entering the finallyBlock production.
	EnterFinallyBlock(c *FinallyBlockContext)

	// EnterJumpExpression is called when entering the jumpExpression production.
	EnterJumpExpression(c *JumpExpressionContext)

	// EnterCallableReference is called when entering the callableReference production.
	EnterCallableReference(c *CallableReferenceContext)

	// EnterAssignmentAndOperator is called when entering the assignmentAndOperator production.
	EnterAssignmentAndOperator(c *AssignmentAndOperatorContext)

	// EnterEqualityOperator is called when entering the equalityOperator production.
	EnterEqualityOperator(c *EqualityOperatorContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterInOperator is called when entering the inOperator production.
	EnterInOperator(c *InOperatorContext)

	// EnterIsOperator is called when entering the isOperator production.
	EnterIsOperator(c *IsOperatorContext)

	// EnterAdditiveOperator is called when entering the additiveOperator production.
	EnterAdditiveOperator(c *AdditiveOperatorContext)

	// EnterMultiplicativeOperator is called when entering the multiplicativeOperator production.
	EnterMultiplicativeOperator(c *MultiplicativeOperatorContext)

	// EnterAsOperator is called when entering the asOperator production.
	EnterAsOperator(c *AsOperatorContext)

	// EnterPrefixUnaryOperator is called when entering the prefixUnaryOperator production.
	EnterPrefixUnaryOperator(c *PrefixUnaryOperatorContext)

	// EnterPostfixUnaryOperator is called when entering the postfixUnaryOperator production.
	EnterPostfixUnaryOperator(c *PostfixUnaryOperatorContext)

	// EnterExcl is called when entering the excl production.
	EnterExcl(c *ExclContext)

	// EnterMemberAccessOperator is called when entering the memberAccessOperator production.
	EnterMemberAccessOperator(c *MemberAccessOperatorContext)

	// EnterSafeNav is called when entering the safeNav production.
	EnterSafeNav(c *SafeNavContext)

	// EnterModifiers is called when entering the modifiers production.
	EnterModifiers(c *ModifiersContext)

	// EnterParameterModifiers is called when entering the parameterModifiers production.
	EnterParameterModifiers(c *ParameterModifiersContext)

	// EnterModifier is called when entering the modifier production.
	EnterModifier(c *ModifierContext)

	// EnterTypeModifiers is called when entering the typeModifiers production.
	EnterTypeModifiers(c *TypeModifiersContext)

	// EnterTypeModifier is called when entering the typeModifier production.
	EnterTypeModifier(c *TypeModifierContext)

	// EnterClassModifier is called when entering the classModifier production.
	EnterClassModifier(c *ClassModifierContext)

	// EnterMemberModifier is called when entering the memberModifier production.
	EnterMemberModifier(c *MemberModifierContext)

	// EnterVisibilityModifier is called when entering the visibilityModifier production.
	EnterVisibilityModifier(c *VisibilityModifierContext)

	// EnterVarianceModifier is called when entering the varianceModifier production.
	EnterVarianceModifier(c *VarianceModifierContext)

	// EnterTypeParameterModifiers is called when entering the typeParameterModifiers production.
	EnterTypeParameterModifiers(c *TypeParameterModifiersContext)

	// EnterTypeParameterModifier is called when entering the typeParameterModifier production.
	EnterTypeParameterModifier(c *TypeParameterModifierContext)

	// EnterFunctionModifier is called when entering the functionModifier production.
	EnterFunctionModifier(c *FunctionModifierContext)

	// EnterPropertyModifier is called when entering the propertyModifier production.
	EnterPropertyModifier(c *PropertyModifierContext)

	// EnterInheritanceModifier is called when entering the inheritanceModifier production.
	EnterInheritanceModifier(c *InheritanceModifierContext)

	// EnterParameterModifier is called when entering the parameterModifier production.
	EnterParameterModifier(c *ParameterModifierContext)

	// EnterReificationModifier is called when entering the reificationModifier production.
	EnterReificationModifier(c *ReificationModifierContext)

	// EnterPlatformModifier is called when entering the platformModifier production.
	EnterPlatformModifier(c *PlatformModifierContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterSingleAnnotation is called when entering the singleAnnotation production.
	EnterSingleAnnotation(c *SingleAnnotationContext)

	// EnterMultiAnnotation is called when entering the multiAnnotation production.
	EnterMultiAnnotation(c *MultiAnnotationContext)

	// EnterAnnotationUseSiteTarget is called when entering the annotationUseSiteTarget production.
	EnterAnnotationUseSiteTarget(c *AnnotationUseSiteTargetContext)

	// EnterUnescapedAnnotation is called when entering the unescapedAnnotation production.
	EnterUnescapedAnnotation(c *UnescapedAnnotationContext)

	// EnterSimpleIdentifier is called when entering the simpleIdentifier production.
	EnterSimpleIdentifier(c *SimpleIdentifierContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitKotlinFile is called when exiting the kotlinFile production.
	ExitKotlinFile(c *KotlinFileContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)

	// ExitShebangLine is called when exiting the shebangLine production.
	ExitShebangLine(c *ShebangLineContext)

	// ExitFileAnnotation is called when exiting the fileAnnotation production.
	ExitFileAnnotation(c *FileAnnotationContext)

	// ExitPackageHeader is called when exiting the packageHeader production.
	ExitPackageHeader(c *PackageHeaderContext)

	// ExitImportList is called when exiting the importList production.
	ExitImportList(c *ImportListContext)

	// ExitImportHeader is called when exiting the importHeader production.
	ExitImportHeader(c *ImportHeaderContext)

	// ExitImportAlias is called when exiting the importAlias production.
	ExitImportAlias(c *ImportAliasContext)

	// ExitTopLevelObject is called when exiting the topLevelObject production.
	ExitTopLevelObject(c *TopLevelObjectContext)

	// ExitTypeAlias is called when exiting the typeAlias production.
	ExitTypeAlias(c *TypeAliasContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitPrimaryConstructor is called when exiting the primaryConstructor production.
	ExitPrimaryConstructor(c *PrimaryConstructorContext)

	// ExitClassBody is called when exiting the classBody production.
	ExitClassBody(c *ClassBodyContext)

	// ExitClassParameters is called when exiting the classParameters production.
	ExitClassParameters(c *ClassParametersContext)

	// ExitClassParameter is called when exiting the classParameter production.
	ExitClassParameter(c *ClassParameterContext)

	// ExitDelegationSpecifiers is called when exiting the delegationSpecifiers production.
	ExitDelegationSpecifiers(c *DelegationSpecifiersContext)

	// ExitDelegationSpecifier is called when exiting the delegationSpecifier production.
	ExitDelegationSpecifier(c *DelegationSpecifierContext)

	// ExitConstructorInvocation is called when exiting the constructorInvocation production.
	ExitConstructorInvocation(c *ConstructorInvocationContext)

	// ExitAnnotatedDelegationSpecifier is called when exiting the annotatedDelegationSpecifier production.
	ExitAnnotatedDelegationSpecifier(c *AnnotatedDelegationSpecifierContext)

	// ExitExplicitDelegation is called when exiting the explicitDelegation production.
	ExitExplicitDelegation(c *ExplicitDelegationContext)

	// ExitTypeParameters is called when exiting the typeParameters production.
	ExitTypeParameters(c *TypeParametersContext)

	// ExitTypeParameter is called when exiting the typeParameter production.
	ExitTypeParameter(c *TypeParameterContext)

	// ExitTypeConstraints is called when exiting the typeConstraints production.
	ExitTypeConstraints(c *TypeConstraintsContext)

	// ExitTypeConstraint is called when exiting the typeConstraint production.
	ExitTypeConstraint(c *TypeConstraintContext)

	// ExitClassMemberDeclarations is called when exiting the classMemberDeclarations production.
	ExitClassMemberDeclarations(c *ClassMemberDeclarationsContext)

	// ExitClassMemberDeclaration is called when exiting the classMemberDeclaration production.
	ExitClassMemberDeclaration(c *ClassMemberDeclarationContext)

	// ExitAnonymousInitializer is called when exiting the anonymousInitializer production.
	ExitAnonymousInitializer(c *AnonymousInitializerContext)

	// ExitCompanionObject is called when exiting the companionObject production.
	ExitCompanionObject(c *CompanionObjectContext)

	// ExitFunctionValueParameters is called when exiting the functionValueParameters production.
	ExitFunctionValueParameters(c *FunctionValueParametersContext)

	// ExitFunctionValueParameter is called when exiting the functionValueParameter production.
	ExitFunctionValueParameter(c *FunctionValueParameterContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFunctionBody is called when exiting the functionBody production.
	ExitFunctionBody(c *FunctionBodyContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitMultiVariableDeclaration is called when exiting the multiVariableDeclaration production.
	ExitMultiVariableDeclaration(c *MultiVariableDeclarationContext)

	// ExitPropertyDeclaration is called when exiting the propertyDeclaration production.
	ExitPropertyDeclaration(c *PropertyDeclarationContext)

	// ExitPropertyDelegate is called when exiting the propertyDelegate production.
	ExitPropertyDelegate(c *PropertyDelegateContext)

	// ExitGetter is called when exiting the getter production.
	ExitGetter(c *GetterContext)

	// ExitSetter is called when exiting the setter production.
	ExitSetter(c *SetterContext)

	// ExitParametersWithOptionalType is called when exiting the parametersWithOptionalType production.
	ExitParametersWithOptionalType(c *ParametersWithOptionalTypeContext)

	// ExitFunctionValueParameterWithOptionalType is called when exiting the functionValueParameterWithOptionalType production.
	ExitFunctionValueParameterWithOptionalType(c *FunctionValueParameterWithOptionalTypeContext)

	// ExitParameterWithOptionalType is called when exiting the parameterWithOptionalType production.
	ExitParameterWithOptionalType(c *ParameterWithOptionalTypeContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitObjectDeclaration is called when exiting the objectDeclaration production.
	ExitObjectDeclaration(c *ObjectDeclarationContext)

	// ExitSecondaryConstructor is called when exiting the secondaryConstructor production.
	ExitSecondaryConstructor(c *SecondaryConstructorContext)

	// ExitConstructorDelegationCall is called when exiting the constructorDelegationCall production.
	ExitConstructorDelegationCall(c *ConstructorDelegationCallContext)

	// ExitEnumClassBody is called when exiting the enumClassBody production.
	ExitEnumClassBody(c *EnumClassBodyContext)

	// ExitEnumEntries is called when exiting the enumEntries production.
	ExitEnumEntries(c *EnumEntriesContext)

	// ExitEnumEntry is called when exiting the enumEntry production.
	ExitEnumEntry(c *EnumEntryContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitTypeReference is called when exiting the typeReference production.
	ExitTypeReference(c *TypeReferenceContext)

	// ExitNullableType is called when exiting the nullableType production.
	ExitNullableType(c *NullableTypeContext)

	// ExitQuest is called when exiting the quest production.
	ExitQuest(c *QuestContext)

	// ExitUserType is called when exiting the userType production.
	ExitUserType(c *UserTypeContext)

	// ExitSimpleUserType is called when exiting the simpleUserType production.
	ExitSimpleUserType(c *SimpleUserTypeContext)

	// ExitTypeProjection is called when exiting the typeProjection production.
	ExitTypeProjection(c *TypeProjectionContext)

	// ExitTypeProjectionModifiers is called when exiting the typeProjectionModifiers production.
	ExitTypeProjectionModifiers(c *TypeProjectionModifiersContext)

	// ExitTypeProjectionModifier is called when exiting the typeProjectionModifier production.
	ExitTypeProjectionModifier(c *TypeProjectionModifierContext)

	// ExitFunctionType is called when exiting the functionType production.
	ExitFunctionType(c *FunctionTypeContext)

	// ExitFunctionTypeParameters is called when exiting the functionTypeParameters production.
	ExitFunctionTypeParameters(c *FunctionTypeParametersContext)

	// ExitParenthesizedType is called when exiting the parenthesizedType production.
	ExitParenthesizedType(c *ParenthesizedTypeContext)

	// ExitReceiverType is called when exiting the receiverType production.
	ExitReceiverType(c *ReceiverTypeContext)

	// ExitParenthesizedUserType is called when exiting the parenthesizedUserType production.
	ExitParenthesizedUserType(c *ParenthesizedUserTypeContext)

	// ExitDefinitelyNonNullableType is called when exiting the definitelyNonNullableType production.
	ExitDefinitelyNonNullableType(c *DefinitelyNonNullableTypeContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitControlStructureBody is called when exiting the controlStructureBody production.
	ExitControlStructureBody(c *ControlStructureBodyContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitLoopStatement is called when exiting the loopStatement production.
	ExitLoopStatement(c *LoopStatementContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitDoWhileStatement is called when exiting the doWhileStatement production.
	ExitDoWhileStatement(c *DoWhileStatementContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitSemi is called when exiting the semi production.
	ExitSemi(c *SemiContext)

	// ExitSemis is called when exiting the semis production.
	ExitSemis(c *SemisContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitInfixFunctionCall is called when exiting the infixFunctionCall production.
	ExitInfixFunctionCall(c *InfixFunctionCallContext)

	// ExitPrefixUnaryExpression is called when exiting the prefixUnaryExpression production.
	ExitPrefixUnaryExpression(c *PrefixUnaryExpressionContext)

	// ExitUnaryPrefix is called when exiting the unaryPrefix production.
	ExitUnaryPrefix(c *UnaryPrefixContext)

	// ExitPostfixUnaryExpression is called when exiting the postfixUnaryExpression production.
	ExitPostfixUnaryExpression(c *PostfixUnaryExpressionContext)

	// ExitPostfixUnarySuffix is called when exiting the postfixUnarySuffix production.
	ExitPostfixUnarySuffix(c *PostfixUnarySuffixContext)

	// ExitDirectlyAssignableExpression is called when exiting the directlyAssignableExpression production.
	ExitDirectlyAssignableExpression(c *DirectlyAssignableExpressionContext)

	// ExitParenthesizedDirectlyAssignableExpression is called when exiting the parenthesizedDirectlyAssignableExpression production.
	ExitParenthesizedDirectlyAssignableExpression(c *ParenthesizedDirectlyAssignableExpressionContext)

	// ExitAssignableExpression is called when exiting the assignableExpression production.
	ExitAssignableExpression(c *AssignableExpressionContext)

	// ExitParenthesizedAssignableExpression is called when exiting the parenthesizedAssignableExpression production.
	ExitParenthesizedAssignableExpression(c *ParenthesizedAssignableExpressionContext)

	// ExitAssignableSuffix is called when exiting the assignableSuffix production.
	ExitAssignableSuffix(c *AssignableSuffixContext)

	// ExitIndexingSuffix is called when exiting the indexingSuffix production.
	ExitIndexingSuffix(c *IndexingSuffixContext)

	// ExitNavigationSuffix is called when exiting the navigationSuffix production.
	ExitNavigationSuffix(c *NavigationSuffixContext)

	// ExitCallSuffix is called when exiting the callSuffix production.
	ExitCallSuffix(c *CallSuffixContext)

	// ExitAnnotatedLambda is called when exiting the annotatedLambda production.
	ExitAnnotatedLambda(c *AnnotatedLambdaContext)

	// ExitTypeArguments is called when exiting the typeArguments production.
	ExitTypeArguments(c *TypeArgumentsContext)

	// ExitValueArguments is called when exiting the valueArguments production.
	ExitValueArguments(c *ValueArgumentsContext)

	// ExitValueArgument is called when exiting the valueArgument production.
	ExitValueArgument(c *ValueArgumentContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitParenthesizedExpression is called when exiting the parenthesizedExpression production.
	ExitParenthesizedExpression(c *ParenthesizedExpressionContext)

	// ExitCollectionLiteral is called when exiting the collectionLiteral production.
	ExitCollectionLiteral(c *CollectionLiteralContext)

	// ExitLiteralConstant is called when exiting the literalConstant production.
	ExitLiteralConstant(c *LiteralConstantContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitLineStringLiteral is called when exiting the lineStringLiteral production.
	ExitLineStringLiteral(c *LineStringLiteralContext)

	// ExitMultiLineStringLiteral is called when exiting the multiLineStringLiteral production.
	ExitMultiLineStringLiteral(c *MultiLineStringLiteralContext)

	// ExitLineStringContent is called when exiting the lineStringContent production.
	ExitLineStringContent(c *LineStringContentContext)

	// ExitLineStringExpression is called when exiting the lineStringExpression production.
	ExitLineStringExpression(c *LineStringExpressionContext)

	// ExitMultiLineStringContent is called when exiting the multiLineStringContent production.
	ExitMultiLineStringContent(c *MultiLineStringContentContext)

	// ExitMultiLineStringExpression is called when exiting the multiLineStringExpression production.
	ExitMultiLineStringExpression(c *MultiLineStringExpressionContext)

	// ExitLambdaLiteral is called when exiting the lambdaLiteral production.
	ExitLambdaLiteral(c *LambdaLiteralContext)

	// ExitLambdaParameters is called when exiting the lambdaParameters production.
	ExitLambdaParameters(c *LambdaParametersContext)

	// ExitLambdaParameter is called when exiting the lambdaParameter production.
	ExitLambdaParameter(c *LambdaParameterContext)

	// ExitAnonymousFunction is called when exiting the anonymousFunction production.
	ExitAnonymousFunction(c *AnonymousFunctionContext)

	// ExitFunctionLiteral is called when exiting the functionLiteral production.
	ExitFunctionLiteral(c *FunctionLiteralContext)

	// ExitObjectLiteral is called when exiting the objectLiteral production.
	ExitObjectLiteral(c *ObjectLiteralContext)

	// ExitThisExpression is called when exiting the thisExpression production.
	ExitThisExpression(c *ThisExpressionContext)

	// ExitSuperExpression is called when exiting the superExpression production.
	ExitSuperExpression(c *SuperExpressionContext)

	// ExitIfExpression is called when exiting the ifExpression production.
	ExitIfExpression(c *IfExpressionContext)

	// ExitWhenSubject is called when exiting the whenSubject production.
	ExitWhenSubject(c *WhenSubjectContext)

	// ExitWhenExpression is called when exiting the whenExpression production.
	ExitWhenExpression(c *WhenExpressionContext)

	// ExitWhenEntry is called when exiting the whenEntry production.
	ExitWhenEntry(c *WhenEntryContext)

	// ExitWhenCondition is called when exiting the whenCondition production.
	ExitWhenCondition(c *WhenConditionContext)

	// ExitRangeTest is called when exiting the rangeTest production.
	ExitRangeTest(c *RangeTestContext)

	// ExitTypeTest is called when exiting the typeTest production.
	ExitTypeTest(c *TypeTestContext)

	// ExitTryExpression is called when exiting the tryExpression production.
	ExitTryExpression(c *TryExpressionContext)

	// ExitCatchBlock is called when exiting the catchBlock production.
	ExitCatchBlock(c *CatchBlockContext)

	// ExitFinallyBlock is called when exiting the finallyBlock production.
	ExitFinallyBlock(c *FinallyBlockContext)

	// ExitJumpExpression is called when exiting the jumpExpression production.
	ExitJumpExpression(c *JumpExpressionContext)

	// ExitCallableReference is called when exiting the callableReference production.
	ExitCallableReference(c *CallableReferenceContext)

	// ExitAssignmentAndOperator is called when exiting the assignmentAndOperator production.
	ExitAssignmentAndOperator(c *AssignmentAndOperatorContext)

	// ExitEqualityOperator is called when exiting the equalityOperator production.
	ExitEqualityOperator(c *EqualityOperatorContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitInOperator is called when exiting the inOperator production.
	ExitInOperator(c *InOperatorContext)

	// ExitIsOperator is called when exiting the isOperator production.
	ExitIsOperator(c *IsOperatorContext)

	// ExitAdditiveOperator is called when exiting the additiveOperator production.
	ExitAdditiveOperator(c *AdditiveOperatorContext)

	// ExitMultiplicativeOperator is called when exiting the multiplicativeOperator production.
	ExitMultiplicativeOperator(c *MultiplicativeOperatorContext)

	// ExitAsOperator is called when exiting the asOperator production.
	ExitAsOperator(c *AsOperatorContext)

	// ExitPrefixUnaryOperator is called when exiting the prefixUnaryOperator production.
	ExitPrefixUnaryOperator(c *PrefixUnaryOperatorContext)

	// ExitPostfixUnaryOperator is called when exiting the postfixUnaryOperator production.
	ExitPostfixUnaryOperator(c *PostfixUnaryOperatorContext)

	// ExitExcl is called when exiting the excl production.
	ExitExcl(c *ExclContext)

	// ExitMemberAccessOperator is called when exiting the memberAccessOperator production.
	ExitMemberAccessOperator(c *MemberAccessOperatorContext)

	// ExitSafeNav is called when exiting the safeNav production.
	ExitSafeNav(c *SafeNavContext)

	// ExitModifiers is called when exiting the modifiers production.
	ExitModifiers(c *ModifiersContext)

	// ExitParameterModifiers is called when exiting the parameterModifiers production.
	ExitParameterModifiers(c *ParameterModifiersContext)

	// ExitModifier is called when exiting the modifier production.
	ExitModifier(c *ModifierContext)

	// ExitTypeModifiers is called when exiting the typeModifiers production.
	ExitTypeModifiers(c *TypeModifiersContext)

	// ExitTypeModifier is called when exiting the typeModifier production.
	ExitTypeModifier(c *TypeModifierContext)

	// ExitClassModifier is called when exiting the classModifier production.
	ExitClassModifier(c *ClassModifierContext)

	// ExitMemberModifier is called when exiting the memberModifier production.
	ExitMemberModifier(c *MemberModifierContext)

	// ExitVisibilityModifier is called when exiting the visibilityModifier production.
	ExitVisibilityModifier(c *VisibilityModifierContext)

	// ExitVarianceModifier is called when exiting the varianceModifier production.
	ExitVarianceModifier(c *VarianceModifierContext)

	// ExitTypeParameterModifiers is called when exiting the typeParameterModifiers production.
	ExitTypeParameterModifiers(c *TypeParameterModifiersContext)

	// ExitTypeParameterModifier is called when exiting the typeParameterModifier production.
	ExitTypeParameterModifier(c *TypeParameterModifierContext)

	// ExitFunctionModifier is called when exiting the functionModifier production.
	ExitFunctionModifier(c *FunctionModifierContext)

	// ExitPropertyModifier is called when exiting the propertyModifier production.
	ExitPropertyModifier(c *PropertyModifierContext)

	// ExitInheritanceModifier is called when exiting the inheritanceModifier production.
	ExitInheritanceModifier(c *InheritanceModifierContext)

	// ExitParameterModifier is called when exiting the parameterModifier production.
	ExitParameterModifier(c *ParameterModifierContext)

	// ExitReificationModifier is called when exiting the reificationModifier production.
	ExitReificationModifier(c *ReificationModifierContext)

	// ExitPlatformModifier is called when exiting the platformModifier production.
	ExitPlatformModifier(c *PlatformModifierContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitSingleAnnotation is called when exiting the singleAnnotation production.
	ExitSingleAnnotation(c *SingleAnnotationContext)

	// ExitMultiAnnotation is called when exiting the multiAnnotation production.
	ExitMultiAnnotation(c *MultiAnnotationContext)

	// ExitAnnotationUseSiteTarget is called when exiting the annotationUseSiteTarget production.
	ExitAnnotationUseSiteTarget(c *AnnotationUseSiteTargetContext)

	// ExitUnescapedAnnotation is called when exiting the unescapedAnnotation production.
	ExitUnescapedAnnotation(c *UnescapedAnnotationContext)

	// ExitSimpleIdentifier is called when exiting the simpleIdentifier production.
	ExitSimpleIdentifier(c *SimpleIdentifierContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
