#!/usr/bin/env bash
#
# @date October 23
# @author dxb The Space Cowboy
#
####
#
# DESCRIPTION:
# mit-license implied
#
# Feeds the main loop, cycling through strategies and betting strategies.
#
# Creates, in essence, enough data by running enouch cycles through the
# looper.#
# For us to evaluate over time and see differences.
#
# If any.
#
######
LOOPER_BIN="./script/loop_program.sh"

cd $(dirname $0)/..


echo "[" $(date) "] Executing: $0"
#print " Executing: "

SLEEPTIME=6

# all available strategies to test
STRATEGIES=()

# auto-generated single strategies
# STRATEGIES+=("almost_always_hit")
STRATEGIES+=("bi_singlestrat_0")
STRATEGIES+=("bi_singlestrat_1")
STRATEGIES+=("bi_singlestrat_2")
STRATEGIES+=("bi_singlestrat_3")
STRATEGIES+=("bi_singlestrat_19044337011466598084822431930153890437606200350598799181000391436239284335758259504900")
STRATEGIES+=("bi_singlestrat_28566505517199897127233647895230835656409300525898198771500587154358926503637389257350")
STRATEGIES+=("bi_singlestrat_38088674022933196169644863860307780875212400701197598362000782872478568671516519009800")
STRATEGIES+=("bi_singlestrat_47610842528666495212056079825384726094015500876496997952500978590598210839395648762250")
STRATEGIES+=("bi_singlestrat_57133011034399794254467295790461671312818601051796397543001174308717853007274778514700")
STRATEGIES+=("bi_singlestrat_66655179540133093296878511755538616531621701227095797133501370026837495175153908267150")
STRATEGIES+=("bi_singlestrat_76177348045866392339289727720615561750424801402395196724001565744957137343033038019600")
STRATEGIES+=("bi_singlestrat_9522168505733299042411215965076945218803100175299399590500195718119642167879129752450")

# stock strategies
#STRATEGIES+=("almost_always_hit")
#STRATEGIES+=("always_hit")
#STRATEGIES+=("advanced")
#STRATEGIES+=("aggressive")
#STRATEGIES+=("passive")
#STRATEGIES+=("wizard_simple")
#STRATEGIES+=("always_stand")
#STRATEGIES+=("no_bust")
#STRATEGIES+=("no_bust2")


#STRATEGIES+=("autostrat_100000")
#STRATEGIES+=("autostrat_200000")

######### RUN 1
#
#STRATEGIES+=("autostrat_1000000")
#STRATEGIES+=("autostrat_2000000")
#STRATEGIES+=("autostrat_3000000")
#STRATEGIES+=("autostrat_4000000")
#STRATEGIES+=("autostrat_5000000")
#STRATEGIES+=("autostrat_6000000")
#STRATEGIES+=("autostrat_7000000")
#STRATEGIES+=("autostrat_8000000")
#STRATEGIES+=("autostrat_9000000")




#STRATEGIES+=("autostrat_100000000")
#STRATEGIES+=("autostrat_200000000")
#STRATEGIES+=("autostrat_300000000")
#STRATEGIES+=("autostrat_400000000")
#STRATEGIES+=("autostrat_500000000")
#STRATEGIES+=("autostrat_600000000")
#STRATEGIES+=("autostrat_700000000")
#STRATEGIES+=("autostrat_800000000")
#STRATEGIES+=("autostrat_900000000")






# @TODO - loop through betting strategies
#BETTING_STRATEGIES=("bet_streaks" "bet_flat" "bet_break1" "bet_break2")
#BETTING_STRATEGIES=("bet_streaks" "bet_flat")
BETTING_STRATEGIES=("bet_flat")
for strategy in ${STRATEGIES[@]}
do
	echo "[" `date` "]Run strategy: ${strategy}"
	for betting_strategy in ${BETTING_STRATEGIES[@]}
	do
		echo "[" `date` "]Run strategy: ${strategy} agasint betting_strategy: ${betting_strategy}"
		./${LOOPER_BIN} ${strategy} ${betting_strategy}	
		echo "[" `date` "]Finished running strategy: ${strategy} agasint betting_strategy: ${betting_strategy}"
		
	done

	echo "[" `date` "]Finished running ${strategy}"
	sleep ${SLEEPTIME}
done	
	
	
print " END OF LINE: "